FROM golang:1.25-alpine AS builder

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd

FROM scratch

WORKDIR /

COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]