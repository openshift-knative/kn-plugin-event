#!/usr/bin/env bash

set -Eeuo pipefail

repodir="$(dirname "$(dirname "$(readlink -f "${BASH_SOURCE[0]:-$0}")")")"
export BUILD_NUMBER=${BUILD_NUMBER:-$(head -c 128 < /dev/urandom | base64 | fold -w 8 | head -n 1)}
export ARTIFACT_DIR="${ARTIFACT_DIR:-$(dirname "$(mktemp -d -u)")/build-${BUILD_NUMBER}}"
export ARTIFACTS="${ARTIFACTS:-${ARTIFACT_DIR}}/kn-event/e2e-tests"
mkdir -p "${ARTIFACTS}"

# shellcheck disable=SC1090
source "$(go run knative.dev/hack/cmd/script e2e-tests.sh)"

set -Eeuo pipefail

# Default the Wathola home to "/" on Openshift (user ID: 65532)
export KN_PLUGIN_EVENT_WATHOLA_HOMEDIR="${KN_PLUGIN_EVENT_WATHOLA_HOMEDIR:-/}"

if [[ "${KN_PLUGIN_EVENT_INSTALL_SERVERLESS:-true}" == "true" ]]; then
  echo '=== Installing Serverless'
  kubectl apply -f "${repodir}/openshift/deploy/serverless-subscription.yaml"
  wait_until_pods_running openshift-serverless

  kubectl apply \
    -f "${repodir}/openshift/deploy/knative-serving.yaml" \
    -f "${repodir}/openshift/deploy/knative-eventing.yaml"
  kubectl wait --for=condition=Ready --timeout=5m \
    knativeserving knative-serving -n knative-serving
  kubectl wait --for=condition=Ready --timeout=5m \
    knativeeventing knative-eventing -n knative-eventing
fi

if [ -z "${KN_PLUGIN_EVENT_EXECUTABLE:-}" ]; then
  echo '=== Building kn-event'
  go build -ldflags \
    "-X knative.dev/kn-plugin-event/pkg/metadata.Version=$(git describe --tags --dirty --always)
    -X knative.dev/kn-plugin-event/pkg/metadata.Image=${KN_PLUGIN_EVENT_SENDER}" \
    -o "${repodir}/build/_output/bin/kn-event-$(go env GOOS)-$(go env GOARCH)" \
    "${repodir}/cmd/kn-event"
fi

echo '=== Running e2e tests'
go_test_e2e -timeout 10m \
  ./test/pkg/... \
  ./test/images/... \
  || fail_test 'kn-event e2e tests'
go_test_e2e -timeout 10m ./test/e2e/... \
  --images.producer.file="${REPO_ROOT_DIR}/openshift/images.yaml" \
  || fail_test 'kn-event e2e tests'

success
