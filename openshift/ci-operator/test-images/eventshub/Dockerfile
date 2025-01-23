ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime
# Reset the goflags to avoid the -mod=vendor flag
ENV GOFLAGS=''

RUN go build -tags strictfipsruntime -o /usr/bin/main knative.dev/reconciler-test/cmd/eventshub

FROM $GO_RUNTIME

ARG VERSION=knative-v1.16

COPY --from=builder /usr/bin/main /usr/bin/eventshub

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-kn-plugin-event-eventshub-rhel8-container" \
      name="openshift-serverless-1-tech-preview/kn-plugin-event-eventshub-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Kn Plugin Event Eventshub" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Kn Plugin Event Eventshub" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Kn Plugin Event Eventshub" \
      io.k8s.description="Red Hat OpenShift Serverless Kn Plugin Event Eventshub" \
      io.openshift.tags="eventshub"

ENTRYPOINT ["/usr/bin/eventshub"]
