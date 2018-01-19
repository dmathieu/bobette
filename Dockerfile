FROM golang:1.9.2 as build

WORKDIR /go/src/github.com/dmathieu/bobette
COPY . .
RUN go get ./...
RUN go install ./...

FROM ubuntu

ENV DOCKER_VERSION="17.09.0~ce-0~ubuntu" \
    DIND_COMMIT="3b5fac462d21ca164b3778647420016315289034"

RUN apt-get update \
      && apt-get install -y --no-install-recommends \
        curl software-properties-common python-software-properties apt-transport-https e2fsprogs iptables xfsprogs xz-utils \
      && curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - \
      && add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" \
      && apt-get update \
      && addgroup dockremap \
      && useradd -g dockremap dockremap \
      && echo 'dockremap:165536:65536' >> /etc/subuid \
      && echo 'dockremap:165536:65536' >> /etc/subgid \
      && curl "https://raw.githubusercontent.com/docker/docker/${DIND_COMMIT}/hack/dind" -o /usr/local/bin/dind \
      && chmod +x /usr/local/bin/dind \
      && apt-get install -y docker-ce=$DOCKER_VERSION \
      && rm -rf /var/lib/apt/lists/*
VOLUME /var/lib/docker
COPY bin/dockerd-entrypoint /usr/local/bin/
ENTRYPOINT ["dockerd-entrypoint"]

COPY --from=build /go/bin/* /usr/local/bin/
COPY bin bin
CMD builder
