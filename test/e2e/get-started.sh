#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# Configure variables.
NAMESPACE="databend-system"
TIMEOUT="5m"

print_cluster_info() {
  kubectl version
  kubectl cluster-info
  kubectl get nodes
  kubectl get pods -n ${NAMESPACE}
  kubectl describe pod -n ${NAMESPACE}
}

echo "Deploying Minio and wait it to be ready"
kubectl apply -f ./examples/get-started/minio.yaml
(kubectl wait pods --for=condition=ready -n ${NAMESPACE} --timeout ${TIMEOUT} --all) ||
  (
    echo "Failed to wait until Minio is ready" &&
      kubectl get pods -n ${NAMESPACE} &&
      kubectl describe pods -n ${NAMESPACE} &&
      exit 1
  )
print_cluster_info

echo "Creating Tenant"
kubectl apply -f ./examples/get-started/tenant.yaml
(kubectl wait tenant/test --for=condition=created  -n ${NAMESPACE} --timeout ${TIMEOUT}) ||
  (
    echo "Failed to wait until Tenant is created" &&
      kubectl get tenants -n ${NAMESPACE} &&
      kubectl describe tenants -n ${NAMESPACE} &&
      exit 1
  )
print_cluster_info

echo "Creating Warehouse"
kubectl apply -f ./examples/get-started/warehouse.yaml
(kubectl wait pods --for=condition=ready -n ${NAMESPACE} --timeout ${TIMEOUT} --all) ||
  (
    echo "Failed to wait until Warehouse is ready" &&
      kubectl get pods -n ${NAMESPACE} &&
      kubectl describe pods -n ${NAMESPACE} &&
      exit 1
  )
print_cluster_info
