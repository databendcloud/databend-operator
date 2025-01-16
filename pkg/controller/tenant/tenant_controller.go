/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tenant

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/databendcloud/databend-operator/pkg/apis/databendlabs.io/v1alpha1"
	"github.com/databendcloud/databend-operator/pkg/common"
)

type opState int

const (
	creationSucceeded opState = iota
	storageError      opState = iota
	metaError         opState = iota
	builtinUserError  opState = iota
)

const DefaultTimeout = time.Second * 10

// TenantReconciler reconciles a Tenant object
type TenantReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=databendlabs.io,resources=tenants,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=databendlabs.io,resources=tenants/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=databendlabs.io,resources=tenants/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list

func (r *TenantReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx)

	var tenant v1alpha1.Tenant
	if err := r.Get(ctx, req.NamespacedName, &tenant); err != nil {
		if apierrors.IsNotFound(err) {
			log.V(2).Info("Tenant has been deleted", "namespacedName", req.NamespacedName)
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log = log.WithValues("tenant", klog.KObj(&tenant))
	ctx = ctrl.LoggerInto(ctx, log)
	log.V(2).Info("Reconciling Tenant")

	var err error
	originStatus := tenant.Status.DeepCopy()

	// Verify storage configuration
	log.V(5).Info("Verifying storage configurations")
	opState, storageErr := r.verifyStorage(ctx, &tenant)
	setCondition(&tenant, opState)
	if storageErr != nil {
		return ctrl.Result{}, errors.Join(err, r.Status().Update(ctx, &tenant))
	}

	// Verify meta configuration
	log.V(5).Info("Verifying meta configurations")
	opState, metaErr := r.verifyMeta(ctx, &tenant)
	setCondition(&tenant, opState)
	if metaErr != nil {
		return ctrl.Result{}, errors.Join(err, r.Status().Update(ctx, &tenant))
	}

	// Verify built-in users configuration
	log.V(5).Info("Verifying built-in users configurations")
	opState, userErr := r.verifyBuiltinUsers(ctx, &tenant)
	setCondition(&tenant, opState)
	if userErr != nil {
		return ctrl.Result{}, errors.Join(err, r.Status().Update(ctx, &tenant))
	}

	if !equality.Semantic.DeepEqual(&tenant.Status, originStatus) {
		return ctrl.Result{}, r.Status().Update(ctx, &tenant)
	}
	return ctrl.Result{}, err
}

func (r *TenantReconciler) verifyStorage(ctx context.Context, tenant *v1alpha1.Tenant) (opState, error) {
	log := ctrl.LoggerFrom(ctx)

	if tenant.Spec.Storage.S3 == nil {
		return storageError, fmt.Errorf("missing S3 configurations")
	}

	// Get accessKey and secretKey
	s3Config := tenant.Spec.Storage.S3
	var accessKey, secretKey string
	if s3Config.S3Auth.SecretRef != nil {
		log.V(5).Info("Getting credentials from Secret")
		var secret corev1.Secret
		nn := types.NamespacedName{
			Namespace: s3Config.S3Auth.SecretRef.Namespace,
			Name:      s3Config.S3Auth.SecretRef.Name,
		}
		if err := r.Get(ctx, nn, &secret, &client.GetOptions{}); err != nil {
			return storageError, fmt.Errorf("failed to get secret %v", nn)
		}
		accessKey, secretKey = string(secret.Data["accessKey"]), string(secret.Data["secretKey"])
	} else {
		accessKey, secretKey = s3Config.AccessKey, s3Config.SecretKey
	}

	// Test connection to S3
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3Config.Region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(s3Config.Endpoint),
		DisableSSL:  aws.Bool(s3Config.AllowInsecure),
	})
	if err != nil {
		return storageError, fmt.Errorf("failed to create session: %w", err)
	}

	// Check bucket
	svc := s3.New(sess)
	_, err = svc.GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: aws.String(s3Config.BucketName),
	})
	if err != nil {
		return storageError, fmt.Errorf("failed to connect to S3: %w", err)
	}

	return creationSucceeded, nil
}

func (r *TenantReconciler) verifyMeta(ctx context.Context, tenant *v1alpha1.Tenant) (opState, error) {
	log := ctrl.LoggerFrom(ctx)

	if len(tenant.Spec.Meta.Endpoints) == 0 {
		return metaError, fmt.Errorf("missing meta endpoints, at lease 1")
	}

	// Check meta connections
	metaConfig := tenant.Spec.Meta
	for _, endpoint := range metaConfig.Endpoints {
		conn, err := net.DialTimeout("tcp", endpoint, DefaultTimeout)
		if err != nil {
			return metaError, fmt.Errorf("failed to dial to %s, %w", endpoint, err)
		}
		_ = conn.Close()
	}

	// Check secrets
	if metaConfig.MetaAuth.PasswordSecretRef != nil {
		log.V(5).Info("Getting meta password from secret")
		var secret corev1.Secret
		nn := types.NamespacedName{
			Namespace: metaConfig.MetaAuth.PasswordSecretRef.Namespace,
			Name:      metaConfig.MetaAuth.PasswordSecretRef.Name,
		}
		if err := r.Get(ctx, nn, &secret, &client.GetOptions{}); err != nil {
			return storageError, fmt.Errorf("failed to get secret %v", nn)
		}
	}

	return creationSucceeded, nil
}

func (r *TenantReconciler) verifyBuiltinUsers(ctx context.Context, tenant *v1alpha1.Tenant) (opState, error) {
	log := ctrl.LoggerFrom(ctx)

	if len(tenant.Spec.Users) == 0 {
		log.V(5).Info("Use default user")
		return creationSucceeded, nil
	}

	// Check secrets
	for _, user := range tenant.Spec.Users {
		if user.AuthStringSecretRef != nil {
			var secret corev1.Secret
			nn := types.NamespacedName{
				Namespace: user.AuthStringSecretRef.Namespace,
				Name:      user.AuthStringSecretRef.Name,
			}
			if err := r.Get(ctx, nn, &secret, &client.GetOptions{}); err != nil {
				return storageError, fmt.Errorf("failed to get secret %v", nn)
			}
		}
	}

	return creationSucceeded, nil
}

func setCondition(tenant *v1alpha1.Tenant, opState opState) {
	var newCond metav1.Condition
	switch opState {
	case creationSucceeded:
		newCond = metav1.Condition{
			Type:    v1alpha1.TenantCreated,
			Status:  metav1.ConditionTrue,
			Message: common.TenantCreationSucceededMessage,
			Reason:  v1alpha1.TenantCreationSucceededReason,
		}
	case storageError:
		newCond = metav1.Condition{
			Type:    v1alpha1.TenantError,
			Status:  metav1.ConditionFalse,
			Message: common.TenantStorageErrorMessage,
			Reason:  v1alpha1.TenantStorageErrorReason,
		}
	case metaError:
		newCond = metav1.Condition{
			Type:    v1alpha1.TenantError,
			Status:  metav1.ConditionFalse,
			Message: common.TenantMetaErrorMessage,
			Reason:  v1alpha1.TenantMetaErrorReason,
		}
	case builtinUserError:
		newCond = metav1.Condition{
			Type:    v1alpha1.TenantError,
			Status:  metav1.ConditionFalse,
			Message: common.TenantUserErrorMessage,
			Reason:  v1alpha1.TenantUserErrorReason,
		}
	default:
		return
	}
	meta.SetStatusCondition(&tenant.Status.Conditions, newCond)
}

// SetupWithManager sets up the controller with the Manager.
func (r *TenantReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Tenant{}).
		Named("tenant").
		Complete(r)
}
