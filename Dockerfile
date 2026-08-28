##################
# TOOLS (DEBIAN) #
##################

# Use golang (debian) as the base image
FROM golang:1.27-trixie as tools

# Install node.js
COPY --from=node:26-trixie-slim /usr/local /usr/local

# Install golangci-lint
COPY --from=golangci/golangci-lint:v2.13.1 /usr/bin/golangci-lint /usr/local/bin/golangci-lint

# Install sqlc
COPY --from=sqlc/sqlc:1.31.1 /workspace/sqlc /usr/local/bin/sqlc

# Install task
# Change to the official image when it's released
# https://github.com/go-task/task/issues/1801
COPY --from=ghcr.io/mirceanton/taskfile:3.53.1 /usr/local/bin/task /usr/local/bin/task

# Install goose
# Change to the official image when it's released
# https://github.com/pressly/goose/issues/1093
COPY --from=ghcr.io/kukymbr/goose-docker-cmd:3.27.2 /bin/goose /usr/local/bin/goose

# Set build time variables
ARG TARGETPLATFORM=unknown \
    DEBIAN_FRONTEND=noninteractive

# Set environment variables
ENV \
    CGO_ENABLED=0 \
    PIP_BREAK_SYSTEM_PACKAGES=1 \
    TZ=Etc/UTC

# Install system dependencies and prepare system
RUN set -e && \
    # Add PostgreSQL APT repository
    # https://www.postgresql.org/download/linux/debian/
    apt-get update -qq && apt-get install -yqq postgresql-common && \
    /usr/share/postgresql-common/pgdg/apt.postgresql.org.sh -y && \
    # Install APT packages
    apt-get update -qq && apt-get install -yqq \
    curl wget git zip unzip p7zip-full tree tzdata \
    ripgrep ca-certificates python3 python3-pip \
    postgresql-client-13 postgresql-client-14 \
    postgresql-client-15 postgresql-client-16 \
    postgresql-client-17 postgresql-client-18 && \
    rm -rf /var/lib/apt/lists/* && \
    # Git config
    git config --global --add safe.directory '*' && \
    # Create backups directory
    mkdir /backups && \
    chmod 777 /backups

WORKDIR /workspaces/pgbackweb

################
# DEVCONTAINER #
################

FROM tools AS devcontainer

CMD ["sleep", "infinity"]

###########
# BUILDER #
###########

FROM tools AS builder

# Copy and install node dependencies
COPY package.json package-lock.json ./
RUN npm ci

# Copy and install go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the files and build
COPY . .
RUN task build

#######################
# PRODUCTION (DEBIAN) #
#######################

FROM debian:trixie-slim AS production

# Set build time variables
ARG DEBIAN_FRONTEND=noninteractive

# Set environment variables
ENV TZ=Etc/UTC

# Install system dependencies and prepare system
RUN set -e && \
    # Add PostgreSQL APT repository
    # https://www.postgresql.org/download/linux/debian/
    apt-get update -qq && apt-get install -yqq postgresql-common && \
    /usr/share/postgresql-common/pgdg/apt.postgresql.org.sh -y && \
    # Install APT packages
    apt-get update -qq && apt-get install -yqq \
    curl wget zip unzip p7zip-full tzdata ca-certificates \
    postgresql-client-13 postgresql-client-14 \
    postgresql-client-15 postgresql-client-16 \
    postgresql-client-17 postgresql-client-18 && \
    rm -rf /var/lib/apt/lists/* && \
    # Create group and user
    groupadd -g 65532 nonroot && \
    useradd -u 65532 -g nonroot -m -s /bin/sh nonroot && \
    # Create backups directory
    mkdir -p /backups && \
    chown -R nonroot:nonroot /backups

WORKDIR /home/nonroot

COPY --from=builder --chown=nonroot:nonroot /workspaces/pgbackweb/dist/pgbackweb /usr/local/bin/pgbackweb

USER nonroot

EXPOSE 8085

ENTRYPOINT ["/usr/local/bin/pgbackweb"]
