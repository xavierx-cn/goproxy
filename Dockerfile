FROM golang:1.16 as builder

WORKDIR /goproxy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o goproxy

FROM alpine:latest

WORKDIR /goproxy
COPY --from=builder /goproxy .
EXPOSE 8765

CMD ["./goproxy"]


