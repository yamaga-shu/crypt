FROM mcr.microsoft.com/devcontainers/go:1.23

WORKDIR /workspace

RUN apt update && apt upgrade -y openssl

VOLUME [ "/workspace" ]
