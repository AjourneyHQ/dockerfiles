ARG RUBY_VERSION=3.3.4
FROM registry.docker.com/library/ruby:$RUBY_VERSION-slim as base

# Rails app lives here
WORKDIR /rails

# Set production environment
ENV BUNDLE_DEPLOYMENT="1" \
    BUNDLE_PATH="/usr/local/bundle" \
    BUNDLE_WITHOUT="development test" \
    NODE_ENV="production" \
    PATH=/usr/local/node/bin:$PATH \
    PLAYWRIGHT_SKIP_BROWSER_DOWNLOAD=1 \
    RAILS_ENV="production"

# Install packages needed for deployment
RUN apt update -qq && \
    apt install --no-install-recommends -y curl libjemalloc2 libpq-dev libvips && \
    rm -rf /var/lib/apt/lists /var/cache/apt/archives

# Install NodeJS for SSR
ARG NODE_VERSION=22.3.0
RUN curl -sL https://github.com/nodenv/node-build/archive/master.tar.gz | tar xz -C /tmp/ && \
    /tmp/node-build-master/bin/node-build "${NODE_VERSION}" /usr/local/node
