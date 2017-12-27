FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir -p /usr/src/app-backend
COPY . /usr/src/app-backend

WORKDIR /usr/src/app-backend

RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /usr/src/app-backend/app-backend .
EXPOSE 3000
CMD ["./app-backend"]
