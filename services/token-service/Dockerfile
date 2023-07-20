FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY ./pkg/ ./pkg/
COPY ./services/token-service/ ./services/token-service/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /token ./services/token-service/cmd/main.go

FROM alpine:3.18.0
COPY --from=builder /token /token
CMD [ "/tokens" ]
