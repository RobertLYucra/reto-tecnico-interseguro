FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
# Download dependencies
# Since we created go.mod manually and didn't run go mod tidy, we do it here.
# But we need the source code first to know what to tidy, or add imports.
# We will copy everything and run tidy.
COPY . .
RUN go mod tidy
RUN go build -o main cmd/api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
