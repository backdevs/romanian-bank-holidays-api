# [Romanian bank holidays API](https://romanian-bank-holidays.backdevs.net)
A simple Go API that serves the Romanian bank holidays

# Prerequisites
* [Docker](https://docs.docker.com/get-docker/) `^20.10`

# Usage
This API is publicly available at: https://romanian-bank-holidays.backdevs.net

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