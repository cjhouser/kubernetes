FROM golang:1.17-alpine AS builder
WORKDIR /go/src/github.com/cjhouser/counter
COPY counter.go .
RUN go mod init github.com/cjhouser/counter
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o build/counter .

FROM scratch
COPY --from=builder /go/src/github.com/cjhouser/counter/build/counter /
EXPOSE 8080
CMD ["./counter"]

