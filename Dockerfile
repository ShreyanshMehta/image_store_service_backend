FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 4000

RUN go build -o /docker-gs-ping

CMD ["/docker-gs-ping"]