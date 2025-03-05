#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

docker build . -t ${OPERATOR_CI_IMAGE} -f ./Dockerfile
