FROM golang:alpine

WORKDIR /app

COPY api-service/go.mod .
COPY api-service/go.sum .
RUN go mod download

COPY ./api-service .

RUN go build -o /goapp ./cmd/app

EXPOSE 8000

CMD ["/goapp"]