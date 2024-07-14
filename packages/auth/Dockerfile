FROM golang:1.22-alpine

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 80

CMD [ "/go/src/app/main" ]