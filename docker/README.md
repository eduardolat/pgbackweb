# Docker

This folder stores all Dockerfiles required for the project. It is imperative
that both images have the same dependencies and versions for consistency and
reproducibility across environments.

## Dockerfile

The Dockerfile is used for building the production image that will be published
for production environments. It should only contain what is strictly required to
run the application.

## Dockerfile.dev

The Dockerfile.dev is used for building the development (e.g., devcontainers)
and CI environment image. It includes all dependencies included in Dockerfile
and others needed for development and testing.
