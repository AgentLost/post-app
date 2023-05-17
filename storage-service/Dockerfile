FROM golang:alpine

WORKDIR /app

COPY storage-service/go.mod .
COPY storage-service/go.sum .
RUN go mod download

COPY ./storage-service .

RUN go build -o /goapp ./cmd/app

EXPOSE 7000

CMD ["/goapp"]







