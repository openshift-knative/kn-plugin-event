# DO NOT EDIT! Generated Dockerfile for cmd/kn-event.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime
ENV GOFLAGS=''

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/kn-event

FROM $GO_RUNTIME

ARG VERSION=knative-next

COPY --from=builder /usr/bin/main /usr/bin/kn-event
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-kn-plugin-event-kn-event-rhel8-container" \
      name="openshift-serverless-1/kn-plugin-event-kn-event-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Kn Plugin Event Kn Event" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Kn Plugin Event Kn Event" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Kn Plugin Event Kn Event" \
      io.k8s.description="Red Hat OpenShift Serverless Kn Plugin Event Kn Event" \
      io.openshift.tags="kn-event"

ENTRYPOINT ["/usr/bin/kn-event"]
