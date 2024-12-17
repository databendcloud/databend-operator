#!/bin/bash

# This shell is used to auto generate some useful tools for k8s, such as clientset, lister, informer and so on.
# We don't use this tool to generate deepcopy because kubebuilder (controller-tools) has covered that part.

set -o errexit
set -o nounset
set -o pipefail

CURRENT_DIR=$(dirname "${BASH_SOURCE[0]}")
DATABEND_OPERATOR_ROOT=$(realpath "${CURRENT_DIR}/..")
DATABEND_OPERATOR_PKG="github.com/databendcloud/databend-operator"

cd "$CURRENT_DIR/.."

# Get the code-generator binary.
CODEGEN_PKG=$(go list -m -mod=readonly -f "{{.Dir}}" k8s.io/code-generator)
source "${CODEGEN_PKG}/kube_codegen.sh"
echo ">> Using ${CODEGEN_PKG}"

# Generate clients for Databend Operator v1alpha1.
echo "Generating clients for v1alpha1"
kube::codegen::gen_client \
  --boilerplate "${DATABEND_OPERATOR_ROOT}/hack/boilerplate.go.txt" \
  --output-dir "${DATABEND_OPERATOR_ROOT}/pkg/client" \
  --output-pkg "${DATABEND_OPERATOR_PKG}/pkg/client" \
  --with-watch \
  --with-applyconfig \
  "${DATABEND_OPERATOR_ROOT}/pkg/apis"

# Get the kube-openapi binary.
OPENAPI_PKG=$(go list -m -mod=readonly -f "{{.Dir}}" k8s.io/kube-openapi)
echo ">> Using ${OPENAPI_PKG}"

# Generating OpenAPI specification for Databend Operator v1alpha1.
echo "Generating OpenAPI specification for databendlabs.io/v1alpha1"
go run ${OPENAPI_PKG}/cmd/openapi-gen \
  --go-header-file "${DATABEND_OPERATOR_ROOT}/hack/boilerplate.go.txt" \
  --output-pkg "${DATABEND_OPERATOR_PKG}/pkg/apis/databendlabs.io/v1alpha1" \
  --output-dir "${DATABEND_OPERATOR_ROOT}/pkg/apis/databendlabs.io/v1alpha1" \
  --output-file "zz_generated.openapi.go" \
  --report-filename "${DATABEND_OPERATOR_ROOT}/hack/violation_exception_v1alpha1.list" \
  "${DATABEND_OPERATOR_ROOT}/pkg/apis/databendlabs.io/v1alpha1"

# Generating OpenAPI Swagger for Training Operator Databend Operator v1alpha1.
echo "Generate OpenAPI Swagger for databendlabs.io/v1alpha1"
go run hack/swagger/main.go > docs/openapi/swagger-v1alpha1.json
