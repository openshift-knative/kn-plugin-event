# DO NOT EDIT! Generated Dockerfile for cmd/kn-event-sender (modified!!).
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-8-release-golang-1.22-openshift-4.17
ARG GO_RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime
# Reset the goflags to avoid the -mod=vendor flag
ENV GOFLAGS=''

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/kn-event-sender

FROM $GO_RUNTIME

ARG VERSION=knative-v1.15

COPY --from=builder /usr/bin/main /usr/bin/kn-event-sender
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-kn-event-sender-rhel8-container" \
      name="openshift-serverless-1-tech-preview/kn-event-sender-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Kn Event Sender" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Kn Event Sender" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Kn Event Sender" \
      io.k8s.description="Red Hat OpenShift Serverless Kn Event Sender" \
      io.openshift.tags="kn-event-sender"

ENTRYPOINT ["/usr/bin/kn-event-sender"]
