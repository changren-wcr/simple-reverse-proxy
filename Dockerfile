## Build
FROM golang:1.18-buster AS build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY *.go ./
RUN go build -o /app/proxy

FROM alpine:latest
RUN apk add gcompat
COPY --from=build /app/proxy /
ENTRYPOINT ["/proxy"]