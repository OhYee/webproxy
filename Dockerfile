FROM golang:1.18 AS builder

COPY . /data
WORKDIR /data
RUN go mod download
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' .

FROM scratch AS prod

ENTRYPOINT [ "/webproxy" ]
EXPOSE 8000
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /etc/ca-certificates  /etc/ca-certificates
COPY --from=builder /etc/ca-certificates.conf  /etc/ca-certificates.conf

COPY --from=builder /data/webproxy /webproxy
