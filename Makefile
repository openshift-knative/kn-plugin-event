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
