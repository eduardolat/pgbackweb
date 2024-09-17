# To make sure we have the node, and golang binaries
FROM node:20.17.0-bookworm AS node
FROM golang:1.23.1-bookworm AS golang

# Set the base image, general environment variables, and move to temp dir
FROM debian:12.7
ENV DEBIAN_FRONTEND=noninteractive
ENV PATH="$PATH:/usr/local/go/bin"
WORKDIR /app/temp

# Copy node binaries
COPY --from=node /usr/local/bin/ /usr/local/bin/
COPY --from=node /usr/local/include/ /usr/local/include/
COPY --from=node /usr/local/lib/ /usr/local/lib/
COPY --from=node /usr/local/share/ /usr/local/share/

# Copy the golang binaries
COPY --from=golang /usr/local/go /usr/local/go

# Add PostgreSQL repository and install system dependencies
# https://www.postgresql.org/download/linux/debian/
RUN apt update && apt install -y postgresql-common && \
    /usr/share/postgresql-common/pgdg/apt.postgresql.org.sh -y && \
    apt update && apt install -y \
        wget p7zip-full tzdata git \
        postgresql-client-13 postgresql-client-14 \
        postgresql-client-15 postgresql-client-16 && \
    rm -rf /var/lib/apt/lists/*

# Install downloadable binaries
RUN set -e && \
    if [ "${TARGETPLATFORM}" = "linux/arm64" ]; then \
        # Install task
        wget --no-verbose https://github.com/go-task/task/releases/download/v3.38.0/task_linux_arm64.tar.gz && \
        tar -xzf task_linux_arm64.tar.gz && \
        mv ./task /usr/local/bin/task && \
        # Install goose
        wget --no-verbose https://github.com/pressly/goose/releases/download/v3.22.0/goose_linux_arm64 && \
        mv ./goose_linux_arm64 /usr/local/bin/goose && \
        # Install sqlc
        wget --no-verbose https://github.com/sqlc-dev/sqlc/releases/download/v1.27.0/sqlc_1.27.0_linux_arm64.tar.gz && \
        tar -xzf sqlc_1.27.0_linux_arm64.tar.gz && \
        mv ./sqlc /usr/local/bin/sqlc && \
        # Install golangci-lint
        wget --no-verbose https://github.com/golangci/golangci-lint/releases/download/v1.60.3/golangci-lint-1.60.3-linux-arm64.tar.gz && \
        tar -xzf golangci-lint-1.60.3-linux-arm64.tar.gz && \
        mv ./golangci-lint-1.60.3-linux-arm64/golangci-lint /usr/local/bin/golangci-lint && \
        # Install air
        wget --no-verbose https://github.com/air-verse/air/releases/download/v1.52.3/air_1.52.3_linux_arm64 && \
        mv ./air_1.52.3_linux_arm64 /usr/local/bin/air; \
    else \
        # Install task
        wget --no-verbose https://github.com/go-task/task/releases/download/v3.38.0/task_linux_amd64.tar.gz && \
        tar -xzf task_linux_amd64.tar.gz && \
        mv ./task /usr/local/bin/task && \
        # Install goose
        wget --no-verbose https://github.com/pressly/goose/releases/download/v3.22.0/goose_linux_x86_64 && \
        mv ./goose_linux_x86_64 /usr/local/bin/goose && \
        # Install sqlc
        wget --no-verbose https://github.com/sqlc-dev/sqlc/releases/download/v1.27.0/sqlc_1.27.0_linux_amd64.tar.gz && \
        tar -xzf sqlc_1.27.0_linux_amd64.tar.gz && \
        mv ./sqlc /usr/local/bin/sqlc && \
        # Install golangci-lint
        wget --no-verbose https://github.com/golangci/golangci-lint/releases/download/v1.60.3/golangci-lint-1.60.3-linux-amd64.tar.gz && \
        tar -xzf golangci-lint-1.60.3-linux-amd64.tar.gz && \
        mv ./golangci-lint-1.60.3-linux-amd64/golangci-lint /usr/local/bin/golangci-lint && \
        # Install air
        wget --no-verbose https://github.com/air-verse/air/releases/download/v1.52.3/air_1.52.3_linux_amd64 && \
        mv ./air_1.52.3_linux_amd64 /usr/local/bin/air; \
    fi && \
    # Make binaries executable
    chmod +x /usr/local/bin/task && \
    chmod +x /usr/local/bin/goose && \
    chmod +x /usr/local/bin/sqlc && \
    chmod +x /usr/local/bin/golangci-lint && \
    chmod +x /usr/local/bin/air

# Go to the app dir, delete the temporary dir and create backups dir
WORKDIR /app
RUN rm -rf /app/temp && \
    mkdir /backups && \
    chmod 777 /backups

##############
# START HERE #
##############

# Copy and install nodejs dependencies
COPY package.json .
RUN npm install

# Copy and install go dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the files
COPY . .

# Build the app
RUN task build

# Copy change-password binary
RUN cp ./dist/change-password /usr/local/bin/change-password && \
    chmod +x /usr/local/bin/change-password

# Run the app
EXPOSE 8085
CMD ["task", "migrate-serve"]
