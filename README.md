# PG Back Web

<img 
  width="70"
  src="internal/view/static/images/logo.png"
/>

🐘 Effortless PostgreSQL backups with a user-friendly web interface! 🌐💾

## Why PG Back Web?

PG Back Web isn't just another backup tool. It's your trusted ally in ensuring the security and availability of your PostgreSQL data:

- 🎯 **Designed for everyone**: From individual developers to teams.
- ⏱️ **Save time**: Automate your backups and forget about manual tasks.

## Features

- 📦 **Intuitive web interface**: Manage your backups with ease, no database expertise required.
- 📅 **Scheduled backups**: Set it and forget it. PG Back Web takes care of the rest.
- 📈 **Backup monitoring**: Visualize the status of your backups with execution logs.
- 📤 **Instant download**: Access your backups when you need them, directly from the web interface.
- 🖥 **Multi-version support**: Compatible with PostgreSQL 13, 14, 15, and 16.
- 📁 **S3 storage**: Add as many S3 buckets as you want for greater flexibility.
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

For the moment you can reset your password by updating the table `users` in the database.

The password should be bcrypt hashed.

## Next steps

- [ ] Password reset command
- [ ] One-click backup restoration
- [ ] API for on-demand backups
- [ ] Advanced alert system and webhooks
- [ ] Automatic health checks

## Join the Community

Got ideas to improve PG Back Web? Contribute to the project! Every suggestion and pull request is welcome.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

💖 **Love PG Back Web?** Give us a ⭐ on GitHub and share the project with your colleagues. Together, we can make PostgreSQL backups more accessible to everyone!