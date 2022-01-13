# creating a Docker build environment for compile.
FROM golang:alpine AS build-env
WORKDIR /app
ADD . /app
RUN go build -o goserver .

# creating the Docker runtime image, using alpine image.
FROM alpine
# install CA-certs for HTTPS
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /app/goserver /app

EXPOSE 8081

ENTRYPOINT ./goserver
