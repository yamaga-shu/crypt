FROM mcr.microsoft.com/devcontainers/go:1.23

WORKDIR /workspace

RUN apt-get update &&\
    apt-get install -y openssl &&\
    rm -rf /var/lib/apt/lists/*

VOLUME [ "/workspace" ]
