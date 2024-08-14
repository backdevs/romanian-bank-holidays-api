FROM golang:1.23-alpine AS build

WORKDIR /src

COPY . .

RUN apk add build-base \
    && go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM scratch

COPY --from=build /src/main /main

EXPOSE 8080

CMD ["/main"]
