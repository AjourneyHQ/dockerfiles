ARG NODEJS_VERSION=22.8.0
FROM node:${NODEJS_VERSION}-bullseye-slim AS base

RUN apt update -qq && \
    apt install --no-install-recommends -y build-essential ca-certificates curl git pkg-config unzip && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/archives

RUN npm install -g pnpm@9.12.2
