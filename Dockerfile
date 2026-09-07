FROM golang:1.26.8-bookworm

RUN mkdir -p /app
COPY . /app
WORKDIR /app
RUN go build -mod=readonly -o yDaemon ./cmd
ENTRYPOINT /app/yDaemon
