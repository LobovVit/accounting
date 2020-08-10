FROM golang:1.14.6 as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY main.go ./
RUN go get github.com/tarantool/go-tarantool
RUN CGO_ENABLED=0 go build -o main .
CMD ["/app/main"]

FROM alpine:latest as my_im
WORKDIR /root/
COPY --from=builder /app/main ./
CMD ["./main"]