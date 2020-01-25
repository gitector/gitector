FROM golang:1.13-alpine

WORKDIR /go/src/app
RUN apk add git
COPY . .
RUN go build \
    && mv gitector /bin/gitector \
    && chmod +x /bin/gitector

ENTRYPOINT ["/bin/gitector"]