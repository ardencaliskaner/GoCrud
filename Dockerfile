FROM golang:1.14.6-alpine3.12 as builder

COPY go.mod go.sum /go/src/GoCrud/
WORKDIR /go/src/GoCrud
RUN go mod download
COPY . /go/src/GoCrud
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/GoCrud GoCrud


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/GoCrud/build/GoCrud /usr/bin/GoCrud

EXPOSE 8080 8080

ENTRYPOINT ["/usr/bin/GoCrud"]
