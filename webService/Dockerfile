FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY migrations/*.sql ./migrations/

RUN go build -o /webService

EXPOSE 8080

CMD [ "/webService" ]