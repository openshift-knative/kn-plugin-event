FROM docker.io/openshift/origin-base

ADD build/_output/bin/kn-event-sender /usr/bin/kn-event-sender
ENTRYPOINT ["/usr/bin/kn-event-sender"]
