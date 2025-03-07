#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# Configure variables.
NAMESPACE="databend-system"
TIMEOUT="5m"
REGISTRY_HUB=k3d-${REGISTRY_NAME}:${REGISTRY_PORT}

# Create a k3d cluster.
echo "Creating k3d registry and cluster"
k3d registry create ${REGISTRY_NAME} --port 0.0.0.0:${REGISTRY_PORT} -i registry:latest
k3d cluster create --config ./scripts/gha/k3d.yaml ${K3D_CLUSTER}
kubectl config use-context k3d-${K3D_CLUSTER}

# Push the databend-operator image to the k3d registry.
echo "Pushing databend-operator image to k3d registry"
echo -n "password" | docker login ${REGISTRY_HUB} --username admin --password-stdin
docker tag ${OPERATOR_CI_IMAGE} ${REGISTRY_HUB}/${OPERATOR_CI_IMAGE}
docker push ${REGISTRY_HUB}/${OPERATOR_CI_IMAGE}

echo "Update databend operator manifest with newly built image"
cd manifests
kustomize edit set image datafuselabs/databend-operator=${REGISTRY_HUB}/${OPERATOR_CI_IMAGE}

echo "Installing databend operator manifests"
kustomize build . | kubectl apply -f -

(kubectl wait deployment/databend-operator --for=condition=available  -n ${NAMESPACE} --timeout ${TIMEOUT} &&
  kubectl wait pods --for=condition=ready -n ${NAMESPACE} --timeout ${TIMEOUT} --all) ||
  (
    echo "Failed to wait until databend operator is ready" &&
      kubectl get pods -n ${NAMESPACE} &&
      kubectl describe pods -n ${NAMESPACE} &&
      exit 1
  )

print_cluster_info() {
  kubectl version
  kubectl cluster-info
  kubectl get nodes
  kubectl get pods -n ${NAMESPACE}
  kubectl describe pod -n ${NAMESPACE}
}

print_cluster_info

echo "Deploying meta cluster with helm"
helm repo add databend https://charts.databend.com
helm install meta databend/databend-meta --namespace ${NAMESPACE}

(kubectl wait pods --for=condition=ready -n ${NAMESPACE} --timeout ${TIMEOUT} --all) ||
  (
    echo "Failed to wait until meta cluster is ready" &&
      kubectl get pods -n ${NAMESPACE} &&
      kubectl describe pods -n ${NAMESPACE} &&
      exit 1
  )

print_cluster_info
