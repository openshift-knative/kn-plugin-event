# DO NOT EDIT! Generated Dockerfile for cmd/kn-event-sender.
ARG GO_BUILDER=registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.25-openshift-4.20
ARG GO_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal

FROM $GO_BUILDER as builder

WORKDIR /workspace
COPY . .

ENV CGO_ENABLED=1
ENV GOEXPERIMENT=strictfipsruntime
ENV GOFLAGS='-mod=mod'
ENV GOFLAGS=''

RUN go build -tags strictfipsruntime -o /usr/bin/main ./cmd/kn-event-sender

FROM $GO_RUNTIME

ARG VERSION=knative-next

COPY --from=builder /usr/bin/main /usr/bin/kn-event-sender
COPY LICENSE /licenses/

USER 65532

LABEL \
      com.redhat.component="openshift-serverless-1-kn-plugin-event-sender-rhel9-container" \
      name="openshift-serverless-1/kn-plugin-event-sender-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift Serverless 1 Kn Plugin Event Kn Event Sender" \
      maintainer="serverless-support@redhat.com" \
      description="Red Hat OpenShift Serverless 1 Kn Plugin Event Kn Event Sender" \
      io.k8s.display-name="Red Hat OpenShift Serverless 1 Kn Plugin Event Kn Event Sender" \
      io.k8s.description="Red Hat OpenShift Serverless Kn Plugin Event Kn Event Sender" \
      io.openshift.tags="kn-event-sender" \
      vendor="Red Hat, Inc." \
      release=$VERSION \
      cpe="cpe:/a:redhat:openshift_serverless:::el9"

ENTRYPOINT ["/usr/bin/kn-event-sender"]
