# To make sure we have the deno and golang binaries
FROM denoland/deno:debian-2.1.9 AS deno
FROM golang:1.23.1-bookworm AS golang

# Set the base image
FROM debian:12.7

# Declare ARG to make it available in the RUN commands
ARG TARGETPLATFORM
RUN echo "Building for ${TARGETPLATFORM}"
RUN if [ "${TARGETPLATFORM}" != "linux/amd64" ] && [ "${TARGETPLATFORM}" != "linux/arm64" ]; then \
        echo "Unsupported architecture: ${TARGETPLATFORM}" && \
        exit 1; \
    fi

# Set the general environment variables, and move to temp dir
ENV DEBIAN_FRONTEND="noninteractive"
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/usr/local/go-bin"
ENV PATH="$PATH:/usr/local/dl-bin"
ENV GOBIN="/usr/local/go-bin"
RUN mkdir -p /app/temp /usr/local/go-bin /usr/local/dl-bin
WORKDIR /app/temp

# Install deno from docker image
COPY --from=deno /usr/bin/deno /usr/local/bin/deno

# Install golang from docker image
COPY --from=golang /usr/local/go /usr/local/go

# Add PostgreSQL repository and install system dependencies
# https://www.postgresql.org/download/linux/debian/
RUN apt update && apt install -y postgresql-common && \
    /usr/share/postgresql-common/pgdg/apt.postgresql.org.sh -y && \
    apt update && apt install -y \
        wget unzip tzdata git \
        postgresql-client-13 postgresql-client-14 \
        postgresql-client-15 postgresql-client-16 \
        postgresql-client-17 && \
    rm -rf /var/lib/apt/lists/*

# Install downloadable binaries
RUN set -e && \
    if [ "${TARGETPLATFORM}" = "linux/arm64" ]; then \
        echo "Downloading arm64 binaries" && \
        # Install task
        wget --no-verbose https://github.com/go-task/task/releases/download/v3.38.0/task_linux_arm64.tar.gz && \
        tar -xzf task_linux_arm64.tar.gz && \
        mv ./task /usr/local/dl-bin/task && \
        # Install goose
        wget --no-verbose https://github.com/pressly/goose/releases/download/v3.22.0/goose_linux_arm64 && \
        mv ./goose_linux_arm64 /usr/local/dl-bin/goose && \
        # Install sqlc
        wget --no-verbose https://github.com/sqlc-dev/sqlc/releases/download/v1.27.0/sqlc_1.27.0_linux_arm64.tar.gz && \
        tar -xzf sqlc_1.27.0_linux_arm64.tar.gz && \
        mv ./sqlc /usr/local/dl-bin/sqlc && \
        # Install golangci-lint
        wget --no-verbose https://github.com/golangci/golangci-lint/releases/download/v1.60.3/golangci-lint-1.60.3-linux-arm64.tar.gz && \
        tar -xzf golangci-lint-1.60.3-linux-arm64.tar.gz && \
        mv ./golangci-lint-1.60.3-linux-arm64/golangci-lint /usr/local/dl-bin/golangci-lint; \
    else \
        echo "Downloading amd64 binaries" && \
        # Install task
        wget --no-verbose https://github.com/go-task/task/releases/download/v3.38.0/task_linux_amd64.tar.gz && \
        tar -xzf task_linux_amd64.tar.gz && \
        mv ./task /usr/local/dl-bin/task && \
        # Install goose
        wget --no-verbose https://github.com/pressly/goose/releases/download/v3.22.0/goose_linux_x86_64 && \
        mv ./goose_linux_x86_64 /usr/local/dl-bin/goose && \
        # Install sqlc
        wget --no-verbose https://github.com/sqlc-dev/sqlc/releases/download/v1.27.0/sqlc_1.27.0_linux_amd64.tar.gz && \
        tar -xzf sqlc_1.27.0_linux_amd64.tar.gz && \
        mv ./sqlc /usr/local/dl-bin/sqlc && \
        # Install golangci-lint
        wget --no-verbose https://github.com/golangci/golangci-lint/releases/download/v1.60.3/golangci-lint-1.60.3-linux-amd64.tar.gz && \
        tar -xzf golangci-lint-1.60.3-linux-amd64.tar.gz && \
        mv ./golangci-lint-1.60.3-linux-amd64/golangci-lint /usr/local/dl-bin/golangci-lint; \
    fi && \
    # Make binaries executable
    chmod +x /usr/local/dl-bin/*

# Default git config
# https://github.com/golangci/golangci-lint/issues/4033
RUN git config --global --add safe.directory '*'

# Go to the app dir, delete the temporary dir and create backups dir
WORKDIR /app
RUN rm -rf /app/temp && \
    mkdir /backups && \
    chmod 777 /backups

##############
# START HERE #
##############

# Copy and install deno dependencies
COPY deno.json .
COPY deno.lock .
RUN deno install

# Copy and install go dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the files
COPY . .

# Fix permissions if needed
RUN task fixperms

# Build the app
RUN task build

# Copy change-password binary
RUN cp ./dist/change-password /usr/local/bin/change-password && \
    chmod +x /usr/local/bin/change-password

# Run the app
EXPOSE 8085
CMD ["task", "migrate-serve"]
