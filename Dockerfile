FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY control-plane/go.mod control-plane/go.sum ./
RUN go mod download
COPY control-plane/main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o fabric-control-plane .
FROM alpine:3.18
RUN apk --no-cache add ca-certificates
RUN addgroup -g 1000 -S appgroup && adduser -u 1000 -S appuser -G appgroup
WORKDIR /app
COPY --from=builder /build/fabric-control-plane .
RUN chown -R appuser:appgroup /app && chmod 755 /app/fabric-control-plane
USER appuser
EXPOSE 8080 9090
ENTRYPOINT ["/app/fabric-control-plane"]
