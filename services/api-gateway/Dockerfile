FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY ./pkg/ ./pkg
COPY ./services/api-gateway/ ./services/api-gateway/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /api-gateway ./services/api-gateway/cmd/main.go

FROM alpine:3.18.0
COPY --from=builder /api-gateway /api-gateway
CMD [ "/api-gateway" ]
