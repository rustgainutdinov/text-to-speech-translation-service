FROM golang:1.17-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main ./cmd
EXPOSE 8080
ADD ./googleConf/top-reef-315512-42fe3ca28b09.json .
CMD ["/app/main"]