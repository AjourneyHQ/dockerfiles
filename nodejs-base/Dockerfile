ARG NODEJS_VERSION=22.8.0
FROM node:${NODEJS_VERSION}-bullseye-slim AS base

RUN apt update -qq && \
    apt install --no-install-recommends -y build-essential ca-certificates curl git pkg-config unzip && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/archives

RUN npm install -g bun@1.1.29

ARG INSTALL_DENO=false
ENV DENO_INSTALL="/usr/local"

RUN if [ "$INSTALL_DENO" = "true" ]; then \
      curl -fsSL https://deno.land/install.sh | sh && \
      deno upgrade --version 2.0.0; \
    fi
