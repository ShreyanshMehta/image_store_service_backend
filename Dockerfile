FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 4000

ENV PORT 4000
ENV DB_HOST localhost
ENV DB_PORT 3306
ENV DB_USER root
ENV DB_PASSWORD 1234
ENV DB_DRIVER mysql
ENV DB_NAME image_store_service

RUN go build -o /docker-backend-service

CMD ["/docker-backend-service"]