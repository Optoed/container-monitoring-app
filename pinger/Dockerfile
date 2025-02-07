FROM golang:1.22 AS build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o pinger .

FROM debian:bookworm-slim

COPY --from=build /app/pinger /usr/local/bin/pinger

RUN apt-get update && apt-get install -y iputils-ping
RUN apt-get update && apt-get install -y docker.io

EXPOSE 9090

CMD ["/usr/local/bin/pinger"]
