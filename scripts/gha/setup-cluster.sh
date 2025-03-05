#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

echo "Kind load newly locally built image"
# use cluster name which is used in github actions kind create
kind load docker-image ${OPERATOR_CI_IMAGE} --name ${KIND_CLUSTER}

echo "Update databend operator manifest with newly built image"
cd manifests
kustomize edit set image datafuselabs/databend-operator=${OPERATOR_CI_IMAGE}

echo "Installing databend operator manifests"
kustomize build . | kubectl apply -f -

TIMEOUT=30
until kubectl get pods -n databend-system | grep databend-operator | grep 1/1 || [[ $TIMEOUT -eq 1 ]] && exit 1; do
  sleep 10
  TIMEOUT=$(( TIMEOUT - 1 ))
done

print_cluster_info() {
  kubectl version
  kubectl cluster-info
  kubectl get nodes
  kubectl get pods -n databend-system
  kubectl describe pod -n databend-system
}

print_cluster_info

echo "Installing Helm"
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

echo "Deploying meta cluster with helm"
helm repo add databend https://charts.databend.com
helm install meta databend/databend-meta --namespace databend-system

TIMEOUT=30
until kubectl get pods -n databend-system | grep databend-meta | grep 1/1 || [[ $TIMEOUT -eq 1 ]] && exit 1; do
  sleep 10
  TIMEOUT=$(( TIMEOUT - 1 ))
done

print_cluster_info
