# Builder
FROM golang:1.22.3-bookworm as builder

WORKDIR /app

COPY go.mod ./go.mod
COPY src/ ./src/

RUN go mod tidy
RUN go build -o ./main ./src

# Runner
FROM debian:stable-slim as runner
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 3000

CMD [ "./main" ]
