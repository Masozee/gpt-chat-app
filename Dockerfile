FROM golang:1.16-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o gpt-chat-app ./cmd/server

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/gpt-chat-app .
COPY --from=builder /app/config.json .

EXPOSE 8080

CMD ["./gpt-chat-app"]