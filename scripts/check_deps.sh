#!/bin/bash

# Function to check if a command runs successfully
check_command() {
  CMD=$1
  DESC=$2

  if $CMD > /dev/null 2>&1; then
    echo "[OK] $DESC"
  else
    echo "[ERROR] $DESC failed"
    exit 1
  fi
}

# Check software from docker images
check_command "go version" "Golang"
check_command "node --version" "Node.js"
check_command "npm --version" "npm"

# Check software installed from apt install
check_command "wget --version" "wget"
check_command "unzip -v" "unzip"
check_command "dpkg -s tzdata" "tzdata"
check_command "git --version" "git"

# Check PostgreSQL clients
check_command "/usr/lib/postgresql/13/bin/psql --version" "PostgreSQL 13 psql"
check_command "/usr/lib/postgresql/13/bin/pg_dump --version" "PostgreSQL 13 pg_dump"
check_command "/usr/lib/postgresql/14/bin/psql --version" "PostgreSQL 14 psql"
check_command "/usr/lib/postgresql/14/bin/pg_dump --version" "PostgreSQL 14 pg_dump"
check_command "/usr/lib/postgresql/15/bin/psql --version" "PostgreSQL 15 psql"
check_command "/usr/lib/postgresql/15/bin/pg_dump --version" "PostgreSQL 15 pg_dump"
check_command "/usr/lib/postgresql/16/bin/psql --version" "PostgreSQL 16 psql"
check_command "/usr/lib/postgresql/16/bin/pg_dump --version" "PostgreSQL 16 pg_dump"
check_command "/usr/lib/postgresql/17/bin/psql --version" "PostgreSQL 17 psql"
check_command "/usr/lib/postgresql/17/bin/pg_dump --version" "PostgreSQL 17 pg_dump"

# Check software installed by downloading binaries
check_command "task --version" "task"
check_command "goose --version" "goose"
check_command "sqlc version" "sqlc"
check_command "golangci-lint --version" "golangci-lint"

echo "All dependencies are working correctly!"
