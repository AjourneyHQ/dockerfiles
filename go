ARG GO_VERSION=1.22.5
FROM golang:${GO_VERSION} as base

WORKDIR /go/src/app

# Set production environment
ENV PATH=/usr/local/node/bin:$PATH \
    PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1

# Install packages needed for deployment
RUN apt update -qq && \
    apt install --no-install-recommends -y curl && \
    rm -rf /var/lib/apt/lists /var/cache/apt/archives

# Install NodeJS for SSR
ARG NODE_VERSION=22.4.1
RUN curl -sL https://github.com/nodenv/node-build/archive/master.tar.gz | tar xz -C /tmp/ && \
    /tmp/node-build-master/bin/node-build "${NODE_VERSION}" /usr/local/node
