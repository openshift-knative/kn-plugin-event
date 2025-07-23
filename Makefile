mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

CI ?= false
PROW_JOB_ID ?= 0
PROW = $(shell [ "$(CI)" = "true" ] && [ "$(PROW_JOB_ID)" != "0" ] && echo "true" || echo "false")

ifeq ($(PROW),true)
	# Openshift CI runs as high UID user
	export HOME = /tmp
	export GOPATH = /tmp/go
	# Reset the goflags to avoid the -mod=vendor flag
	export GOFLAGS =
endif

build:
	./mage build
.PHONY: build

unit:
	./mage test
.PHONY: unit

test: unit
.PHONY: test

e2e:
	openshift/e2e-tests.sh
.PHONY: e2e

dockerfiles:
	go run github.com/openshift-knative/hack/cmd/generate@latest \
		--root-dir "$(current_dir)" \
		--generators dockerfile \
		--app-file-fmt "/usr/bin/%s" \
		--dockerfile-image-builder-fmt "registry.ci.openshift.org/openshift/release:rhel-8-release-golang-%s-openshift-4.17"
.PHONY: dockerfiles
