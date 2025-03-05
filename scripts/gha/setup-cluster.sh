#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# Configure variables.
NAMESPACE="databend-system"
TIMEOUT="5m"

echo "Kind load newly locally built image"
# use cluster name which is used in github actions kind create
kind load docker-image ${OPERATOR_CI_IMAGE} --name ${KIND_CLUSTER}

echo "Update databend operator manifest with newly built image"
cd manifests
kustomize edit set image datafuselabs/databend-operator=${OPERATOR_CI_IMAGE}

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
