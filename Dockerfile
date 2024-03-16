FROM golang:1.21 AS builder
LABEL authors="andrauss"

WORKDIR /app
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server -v cmd/api/main.go
RUN ls -la


FROM scratch
COPY --from=builder /app/server /server

EXPOSE 8080

CMD ["/server"]