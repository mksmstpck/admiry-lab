FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY ./pkg/ ./pkg/
COPY ./services/user-service/ ./services/user-service/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /users ./services/user-service/cmd/main.go

FROM alpine:3.18.0
COPY --from=builder /users /users
CMD [ "/users" ]
