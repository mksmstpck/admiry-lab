FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY ./pkg/ ./pkg/
COPY ./services/company-service/ ./services/company-service/
RUN go mod download
RUN CGO_ENABLED=0 go build -o /companies ./services/company-service/cmd/main.go

FROM alpine:3.18.0
COPY --from=builder /companies /companies
CMD [ "/companies" ]
