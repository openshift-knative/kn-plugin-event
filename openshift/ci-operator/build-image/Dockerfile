# Dockerfile to bootstrap build and test in openshift-ci
FROM registry.ci.openshift.org/ocp/builder:rhel-8-golang-1.17-openshift-4.10

RUN curl -LO https://mirror.openshift.com/pub/openshift-v4/clients/ocp/latest/openshift-client-linux.tar.gz
RUN curl -LO https://mirror.openshift.com/pub/openshift-v4/clients/ocp/latest/sha256sum.txt.gpg
RUN curl -LO https://mirror.openshift.com/pub/openshift-v4/clients/ocp/latest/sha256sum.txt
RUN curl -LO https://mirror.openshift.com/pub/openshift-v4/clients/ocp/latest/release.txt
RUN curl -LO https://www.redhat.com/security/data/fd431d51.txt
RUN gpg --show-keys fd431d51.txt | grep 567E347AD0044ADE55BA8A5F199E2F91FD431D51
RUN gpg --import fd431d51.txt
RUN gpg --list-keys --fingerprint --with-colons | \
      sed -E -n -e 's/^fpr:::::::::([0-9A-F]+):$/\1:6:/p' | \
      gpg --import-ownertrust
RUN gpg --verify sha256sum.txt.gpg
RUN ln -s openshift-client-linux.tar.gz \
    openshift-client-linux-$(cat release.txt | grep Name: | awk '{print $2}').tar.gz
RUN sha256sum --check --ignore-missing sha256sum.txt
RUN tar xzvf openshift-client-linux.tar.gz
RUN install oc /usr/bin
RUN install kubectl /usr/bin
