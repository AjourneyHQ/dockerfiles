FROM registry.docker.com/cayt3r/rails:latest

# Install packages needed to build gems
RUN apt update -qq && \
    apt install --no-install-recommends -y build-essential git pkg-config && \
    rm -rf /var/lib/apt/lists /var/cache/apt/archives

# Install JavaScript dependencies
ARG BUN_VERSION=1.1.20
RUN npm install -g bun@$BUN_VERSION
