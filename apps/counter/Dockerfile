# syntax=docker/dockerfile:1
FROM golang:1.16-alpine AS builder
WORKDIR /go/src/github.com/cjhouser/counter
COPY counter.go .
RUN go mod init github.com/cjhouser/counter
RUN CGO_ENABLED=0 GOOS=linux go build -o build/counter -a -installsuffix cgo .

FROM scratch
COPY --from=builder /go/src/github.com/cjhouser/counter/build/counter /
EXPOSE 8080
CMD ["/counter"]