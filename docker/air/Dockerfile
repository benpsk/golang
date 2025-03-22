FROM golang:1.24

WORKDIR /app

RUN go install github.com/air-verse/air@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download

# RUN go build -v -o /usr/local/bin/app ./...

CMD ["air", "-c", ".air.toml"]

