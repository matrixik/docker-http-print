FROM golang:latest as builder

WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM centurylink/ca-certs
WORKDIR /
COPY --from=builder /go/src/app/app .
ENTRYPOINT ["/app"]

