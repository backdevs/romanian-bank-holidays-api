# [Romanian bank holidays API](https://api.bank-holidays.ro)
A simple Go API that serves the Romanian bank holidays

![Build and Push status](https://github.com/backdevs/romanian-bank-holidays-api/actions/workflows/docker-build-and-push.yml/badge.svg)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/backdevs/romanian-bank-holidays-api)

# Prerequisites
* [Docker](https://docs.docker.com/get-docker/) `^20.10`

# Usage
This API is publicly available at: https://api.bank-holidays.ro/?year=2049

However, if you want, you can run it locally using docker:
```shell
docker run -p 8080:8080 backdevs/romanian-bank-holidays-api
```

And access it at http://localhost:8080/?year=2025


# Build

```shell
docker build \
  --tag backdevs/romanian-bank-holidays-api:1.0.0 \
  --tag backdevs/romanian-bank-holidays-api:latest \
  .
```
