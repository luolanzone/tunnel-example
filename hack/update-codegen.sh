#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# corresponding to go mod init <module>
MODULE=github.com/luolanzone/tunnel-example
# group-version such as foo:v1alpha1
GROUP=gateway
VERSION=v1
GROUP_VERSION=${GROUP}:${VERSION}

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

rm -rf ${MODULE}/pkg/client/{clientset,informers,listers}

$GOPATH/bin/client-gen \
  --clientset-name versioned \
  --input-base "${MODULE}/apis" \
  --input "gateway/v1alpha1" \
  --output-package "${MODULE}/pkg/client/clientset" \
  --go-header-file hack/boilerplate.go.txt

# Generate listers with K8s codegen tools.
$GOPATH/bin/lister-gen \
  --input-dirs "${MODULE}/apis/gateway/v1alpha1" \
  --output-package "${MODULE}/pkg/client/listers" \
  --go-header-file hack/boilerplate.go.txt


# Generate informers with K8s codegen tools.
$GOPATH/bin/informer-gen \
  --input-dirs "${MODULE}/apis/gateway/v1alpha1" \
  --versioned-clientset-package "${MODULE}/pkg/client/clientset/versioned" \
  --listers-package "${MODULE}/pkg/client/listers" \
  --output-package "${MODULE}/pkg/client/informers" \
  --go-header-file hack/boilerplate.go.txt
