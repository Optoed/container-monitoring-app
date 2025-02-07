FROM golang:1.22 AS build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o backend .

FROM debian:bookworm-slim

COPY --from=build /app/backend /usr/local/bin/backend

RUN apt-get update && apt-get install -y libpcap-dev

EXPOSE 8080

CMD ["/usr/local/bin/backend"]
