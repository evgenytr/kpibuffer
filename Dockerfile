FROM golang:1.21
WORKDIR /app
COPY go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /cmd/server/bin/main /cmd/server/main.go
CMD ["bash", "-c", "/cmd/server/bin/main"]