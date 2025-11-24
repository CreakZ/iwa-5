FROM golang:1.24.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

RUN go build -o contacts-app .

EXPOSE 8080

CMD ["./contacts-app"]
