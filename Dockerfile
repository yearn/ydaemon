FROM golang:1.19-bullseye

RUN mkdir -p /app
COPY . /app
WORKDIR /app
RUN go mod tidy
RUN go mod vendor
RUN go build -o yDaemon ./cmd
ENTRYPOINT /app/yDaemon
