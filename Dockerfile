FROM golang:1.22.2-alpine

WORKDIR /app

COPY . .
RUN go build -o jokes cmd/main.go

ENTRYPOINT [ "./jokes" ]

