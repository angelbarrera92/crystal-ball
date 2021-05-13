FROM golang:1.16.3-alpine

WORKDIR /go/src/app
COPY . .

RUN apk add git gcc

RUN go get -d ./...
RUN go install -v ./...

CMD ["crystal-ball"]