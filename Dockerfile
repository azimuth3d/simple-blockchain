FROM golang:1.13.0-alpine as builder
WORKDIR /usr/app/
RUN apk update 
RUN apk --no-cache add git gcc g++ make ca-certificates
RUN go get github.com/cespare/reflex
COPY reflex.conf /
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]