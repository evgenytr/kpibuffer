FROM golang:1.21
WORKDIR /kpibuffer
COPY go.mod ./
COPY . /kpibuffer
RUN CGO_ENABLED=0 GOOS=linux go build -o /kpibuffer/cmd/server/bin/main /kpibuffer/cmd/server/main.go
CMD ["bash", "-c", "/kpibuffer/cmd/server/bin/main"]