FROM golang:1.27-alpine AS builder

WORKDIR /app
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o log-output .

FROM alpine:3.20
COPY --from=builder /app/log-output /log-output

ENTRYPOINT ["/log-output"]