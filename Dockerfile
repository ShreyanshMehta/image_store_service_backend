FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 4000

ENV PORT 4000

RUN go build -o /docker-backend-service

CMD ["/docker-backend-service"]