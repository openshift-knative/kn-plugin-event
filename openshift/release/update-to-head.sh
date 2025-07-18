#!/usr/bin/env bash

set -Eeuo pipefail

project_dir="$(realpath "$(dirname "${BASH_SOURCE[0]:-$0}")/../..")"

cd "$project_dir"

exec go run \
  github.com/openshift-knative/deviate/cmd/deviate@main \
  sync --config "${project_dir}/.deviate.yaml"
