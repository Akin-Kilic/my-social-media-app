FROM golang:1.21.4 AS builder

WORKDIR /src/go

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o netadim main.go

FROM scratch

# take env from build args
ARG VERSION
ENV APP_VERSION=$VERSION

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /bin/netadim

COPY --from=builder /src/go/netadim .

CMD [ "./netadim" ]
