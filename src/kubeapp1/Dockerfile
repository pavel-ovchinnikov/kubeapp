# Stage 1: Build the application
FROM golang:1.24 AS builder
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /kubeapp ./cmd/main.go

# Stage 2: Create a smaller image with the built application
FROM alpine:latest
WORKDIR /
COPY --from=builder /kubeapp /kubeapp
EXPOSE 8890
CMD ["/kubeapp"]
