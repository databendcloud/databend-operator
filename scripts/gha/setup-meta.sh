#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

echo "Deploying meta cluster with helm"
helm repo add databend https://charts.databend.com
helm install meta databend/databend-meta --namespace databend-system

TIMEOUT=30
until kubectl get pods -n databend-system | grep databend-meta | grep 1/1 || [[ $TIMEOUT -eq 1 ]]; do
  sleep 10
  TIMEOUT=$(( TIMEOUT - 1 ))
done

kubectl version
kubectl cluster-info
kubectl get nodes
kubectl get pods -n databend-system
kubectl describe pods -n databend-system
