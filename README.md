<p align="center">
  <h1 align="center">PG Back Web</h1>
  <p align="center">
    <img align="center" width="70" src="https://raw.githubusercontent.com/eduardolat/pgbackweb/main/internal/view/static/images/logo.png"/>
  </p>
  <p align="center">
    🐘 Effortless PostgreSQL backups with a user-friendly web interface! 🌐💾
  </p>
</p>
<p align="center">
  <a href="https://github.com/eduardolat/pgbackweb/actions/workflows/ci.yaml?query=branch%3Amain">
    <img src="https://github.com/eduardolat/pgbackweb/actions/workflows/ci.yaml/badge.svg" alt="CI Status"/>
  </a>
  <a href="https://goreportcard.com/report/eduardolat/pgbackweb">
    <img src="https://goreportcard.com/badge/eduardolat/pgbackweb" alt="Go Report Card"/>
  </a>
  <a href="https://github.com/eduardolat/pgbackweb/releases/latest">
    <img src="https://img.shields.io/github/release/eduardolat/pgbackweb.svg" alt="Release Version"/>
  </a>
  <a href="https://hub.docker.com/r/eduardolat/pgbackweb">
    <img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/eduardolat/pgbackweb"/>
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/github/license/eduardolat/pgbackweb.svg" alt="License"/>
  </a>
  <a href="https://github.com/eduardolat/pgbackweb">
    <img src="https://img.shields.io/github/stars/eduardolat/pgbackweb?style=flat&label=github+stars"/>
  </a>
</p>

## Why PG Back Web?

PG Back Web isn't just another backup tool. It's your trusted ally in ensuring
the security and availability of your PostgreSQL data:

- 🎯 **Designed for everyone**: From individual developers to teams.
- ⏱️ **Save time**: Automate your backups and forget about manual tasks.
- ⚡ **Plug and play**: Don't waste time with complex configurations.

## Features

- 📦 **Intuitive web interface**: Manage your backups with ease, no database
  expertise required.
- 📅 **Scheduled backups**: Set it and forget it. PG Back Web takes care of the
  rest.
- 📈 **Backup monitoring**: Visualize the status of your backups with execution
  logs.
- 📤 **Instant download & restore**: Restore and download your backups when you
  need them, directly from the web interface.
- 🖥 **Multi-version support**: Compatible with PostgreSQL 13, 14, 15, 16,
  17, and 18.
- 📁 **Local & S3 storage**: Store backups locally or add as many S3 buckets as
  you want for greater flexibility.
- ❤️‍🩹 **Health checks**: Automatically check the health of your databases and
  destinations.
- 🔔 **Webhooks**: Get notified when a backup finishes, failed, health check
  fails, or other events.
- 🔒 **Security first**: PGP encryption to protect your sensitive information.
- 🛡️ **Open-source trust**: Open-source code under MIT license, backed by the
  robust pg_dump tool.
- 🌚 **Dark mode**: Because we all love dark mode.

## Installation

PG Back Web is available as a Docker image. You just need to set 3 environment
variables and you're good to go!

Here's an example of how you can run PG Back Web with Docker Compose, feel free
to adapt it to your needs:

```yaml
services:
  pgbackweb:
    image: eduardolat/pgbackweb:latest
    ports:
      - "8085:8085" # Access the web interface at http://localhost:8085
    volumes:
      - ./backups:/backups # If you only use S3 destinations, you don't need this volume
    environment:
      PBW_ENCRYPTION_KEY: "my_secret_key" # Change this to a strong key
      PBW_POSTGRES_CONN_STRING: "postgresql://postgres:password@postgres:5432/pgbackweb?sslmode=disable"
      TZ: "America/Guatemala" # Set your timezone, optional
    depends_on:
      postgres:
        condition: service_healthy

  postgres:
    image: postgres:17
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

You can watch [this youtube video](https://www.youtube.com/watch?v=vf7SLrSO8sw)
to see how easy it is to set up PG Back Web.

## Configuration

You only need to configure the following environment variables:

- `PBW_ENCRYPTION_KEY`: Your encryption key. Generate a strong one and store it
  in a safe place, as PG Back Web uses it to encrypt sensitive data.

- `PBW_POSTGRES_CONN_STRING`: The connection string for the PostgreSQL database
  that will store PG Back Web data.

- `PBW_LISTEN_HOST`: Host for the server to listen on, default 0.0.0.0
  (optional)

- `PBW_LISTEN_PORT`: Port for the server to listen on, default 8085 (optional)

- `TZ`: Your
  [timezone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
  (optional). Default is `UTC`. This impacts logging, backup filenames and
  default timezone in the web interface.

## Screenshot

<img src="https://raw.githubusercontent.com/eduardolat/pgbackweb/main/assets/screenshot.png" />

## Reset password

You can reset your PG Back Web password by running the following command in the
server where PG Back Web is running:

```bash
docker exec -it <container_name_or_id> sh -c change-password
```

You should replace `<container_name_or_id>` with the name or ID of the PG Back
Web container, then just follow the instructions.

## Next steps

In this link you can see a list of features that have been confirmed for future
updates:

<a href="https://github.com/eduardolat/pgbackweb/issues?q=is%3Aissue+is%3Aopen+label%3A%22confirmed+next+step%22">
  Next steps ⏭️
</a>

## Sponsors

🙏 Thank you to the incredible sponsors for supporting this project! Your
contributions help keep PG Back Web running and growing. If you'd like to join
and become a sponsor, please visit the
[sponsorship page](https://buymeacoffee.com/eduardolat) and be part of something
great! 🚀

### 🥇 Gold Sponsors

<table>
  <tr>
    <td align="center">
      <a href="https://buymeacoffee.com/eduardolat">
        <img src="https://raw.githubusercontent.com/eduardolat/pgbackweb/refs/heads/develop/internal/view/static/images/plus-circle.png" height="150" alt="Become a gold sponsor"/>
        <br />
        Become a gold sponsor
      </a>
    </td>
  </tr>
</table>

### 🥈 Silver Sponsors

<table>
  <tr>
    <td align="center">
      <a href="https://fetchgoat.com?utm_source=pgbackweb&utm_medium=referral&utm_campaign=sponsorship">
        <img src="https://raw.githubusercontent.com/eduardolat/pgbackweb/refs/heads/develop/assets/sponsors/FetchGoat.png" height="100" alt="FetchGoat - Simplifying Logistics"/>
        <br />
        FetchGoat - Simplifying Logistics
      </a>
    </td>
    <td align="center">
      <a href="https://buymeacoffee.com/eduardolat">
        <img src="https://raw.githubusercontent.com/eduardolat/pgbackweb/refs/heads/develop/internal/view/static/images/plus-circle.png" height="100" alt="Become a silver sponsor"/>
        <br />
        Become a silver sponsor
      </a>
    </td>
  </tr>
</table>

### 🥉 Bronze Sponsors

<table>
  <tr>
    <td align="center">
      <a href="https://buymeacoffee.com/eduardolat">
        <img src="https://raw.githubusercontent.com/eduardolat/pgbackweb/refs/heads/develop/internal/view/static/images/plus-circle.png" height="80" alt="Become a bronze sponsor"/>
        <br />
        Become a bronze sponsor
      </a>
    </td>
  </tr>
</table>

## Join the Community

Got ideas to improve PG Back Web? Contribute to the project! Every suggestion
and pull request is welcome.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file
for details.

---

💖 **Love PG Back Web?** Give us a ⭐ on GitHub and share the project with your
colleagues. Together, we can make PostgreSQL backups more accessible to
everyone!
