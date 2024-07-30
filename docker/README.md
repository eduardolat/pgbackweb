# Docker

This folder is responsible for storing all Dockerfiles for the different Docker images required for the project. IT IS IMPERATIVE that all images have the same dependencies and versions to ensure consistency and reproducibility across different environments.

To make sure they are identical, please always ensure that all instructions in the Dockerfiles are exactly the same up to the following comment:

```Dockerfile
##############
# START HERE #
##############
```

After the above comment, you can add the necessary instructions for the specific image. Please replicate all dependencies and environment-related instructions in all Dockerfiles.

## Dockerfile

The `Dockerfile` file is the one used for building the production image.

## Dockerfile.dev

The `Dockerfile.dev` file is the one used for building the development environment image.

## Dockerfile.cicd

The `Dockerfile.cicd` file is the one used for building the CI/CD environment image.
