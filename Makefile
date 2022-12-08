CGO_ENABLED=0
GOOS=linux
TEST_IMAGES=./vendor/knative.dev/eventing/test/test_images/wathola-forwarder ./vendor/knative.dev/reconciler-test/cmd/eventshub
TEST_IMAGE_TAG ?= latest

# Target used by github actions.
test-images:
	for img in $(TEST_IMAGES); do \
		KO_DOCKER_REPO=$(DOCKER_REPO_OVERRIDE) ko build --tags=$(TEST_IMAGE_TAG) $(KO_FLAGS) -B $$img || exit $?; \
	done
.PHONY: test-images
