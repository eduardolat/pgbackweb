<p align="center">
  <h1 align="center">PG Back Web</h1>
  <p align="center">
    <img align="center" width="70" src="internal/view/static/images/logo.png"/>
  </p>
  <p align="center">
    🐘 Effortless PostgreSQL backups with a user-friendly web interface! 🌐💾
  </p>
</p>
<p align="center">
  <a href="https://github.com/eduardolat/pgbackweb/actions/workflows/run-all-tests.yaml?query=branch%3Amain">
    <img src="https://github.com/eduardolat/pgbackweb/actions/workflows/run-all-tests.yaml/badge.svg" alt="Tests Status"/>
  </a>
  <a href="https://goreportcard.com/report/eduardolat/pgbackweb">
    <img src="https://goreportcard.com/badge/eduardolat/pgbackweb" alt="Go Report Card"/>
  </a>
  <a href="https://github.com/eduardolat/pgbackweb/releases/latest">
    <img src="https://img.shields.io/github/release/eduardolat/pgbackweb.svg" alt="Release Version"/>
  </a>
  <a href="https://hub.docker.com/r/eduardolat/pgbackweb">
    <img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/eduardolat/pgbackweb"/>
  <a/>
  <a href="LICENSE">
    <img src="https://img.shields.io/github/license/eduardolat/pgbackweb.svg" alt="License"/>
  </a>
</p>

## Why PG Back Web?

PG Back Web isn't just another backup tool. It's your trusted ally in ensuring the security and availability of your PostgreSQL data:

- 🎯 **Designed for everyone**: From individual developers to teams.
- ⏱️ **Save time**: Automate your backups and forget about manual tasks.
- ⚡ **Plug and play**: Don't waste time with complex configurations.

## Features

- 📦 **Intuitive web interface**: Manage your backups with ease, no database expertise required.
- 📅 **Scheduled backups**: Set it and forget it. PG Back Web takes care of the rest.
- 📈 **Backup monitoring**: Visualize the status of your backups with execution logs.
- 📤 **Instant download & restore**: Restore and download your backups when you need them, directly from the web interface.
- 🖥 **Multi-version support**: Compatible with PostgreSQL 13, 14, 15, and 16.
- 📁 **Local & S3 storage**: Store backups locally or add as many S3 buckets as you want for greater flexibility.
- 🔒 **Security first**: PGP encryption to protect your sensitive information.
- 🛡️ **Open-source trust**: Open-source code under MIT license, backed by the robust pg_dump tool.
- 🌚 **Dark mode**: Because we all love dark mode.

## Installation

PG Back Web is available as a Docker image. You just need to set 2 environment variables and you're good to go!

Here's an example of how you can run PG Back Web with Docker Compose, feel free to adapt it to your needs:

```yaml
services:
  pgbackweb:
    image: eduardolat/pgbackweb:latest
    ports:
      - "8085:8085" # Access the web interface at http://localhost:8085
    environment:
      PBW_ENCRYPTION_KEY: "my_secret_key"
      PBW_POSTGRES_CONN_STRING: "postgresql://postgres:password@postgres:5432/pgbackweb?sslmode=disable"
    depends_on:
      postgres:
        condition: service_healthy

  postgres:
    image: postgres:16
    environment:
      POSTGRES_USER: postgres
      POSTGRES_DB: pgbackweb
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - ./data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 5s
      timeout: 5s
      retries: 5
```

## Configuration

You only need to configure the following environment variables:

- `PBW_ENCRYPTION_KEY`: Your encryption key. Generate a strong one and store it in a safe place, as PG Back Web uses it to encrypt sensitive data.

- `PBW_POSTGRES_CONN_STRING`: The connection string for the PostgreSQL database that will store PG Back Web data.

## Screenshots

<img src="screenshots/summary.png" />
<img src="screenshots/backups.png" />
<img src="screenshots/executions.png" />

## Reset password

You can reset your PG Back Web password by running the following command in the server where PG Back Web is running:

```bash
docker exec -it <container_name_or_id> sh -c change-password
```

You should replace `<container_name_or_id>` with the name or ID of the PG Back Web container, then just follow the instructions.

## Next steps

In this link you can see a list of features that have been confirmed for future updates:

<a href="https://github.com/eduardolat/pgbackweb/issues?q=is%3Aissue+is%3Aopen+label%3A%22confirmed+next+step%22">
  Next steps ⏭️
</a>

## Join the Community

Got ideas to improve PG Back Web? Contribute to the project! Every suggestion and pull request is welcome.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

💖 **Love PG Back Web?** Give us a ⭐ on GitHub and share the project with your colleagues. Together, we can make PostgreSQL backups more accessible to everyone!