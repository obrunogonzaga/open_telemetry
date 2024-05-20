FROM golang:1.22-alpine
LABEL authors="bruno gonzaga"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o weatherservice ./cmd/weatherservice

EXPOSE 8081

CMD ["/app/weatherservice"]