FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git ca-certificates && \
    update-ca-certificates && \
    adduser -D -g '' appuser

COPY main.go .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/crud-app

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/crud-app /crud-app
COPY form .
USER 4000
ENTRYPOINT ["/crud-app"]
