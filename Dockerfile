FROM golang:1.18-alpine

WORKDIR /app

COPY ./src/go.mod ./
COPY ./src/go.sum ./
RUN go mod download

COPY ./src ./

WORKDIR /app/cmd
RUN go build -o /api

EXPOSE 8080

CMD ["/api"]