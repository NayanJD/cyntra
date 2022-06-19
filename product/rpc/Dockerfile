FROM golang:1.18.3-alpine3.16 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY product/rpc/etc /app/etc

# Chose arm64 for running the image in M1 chip. Use amd64 for linux/intel chip mac.
RUN wget -O grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.11/grpc_health_probe-linux-arm64

RUN chmod u+x grpc_health_probe

RUN go build -ldflags="-s -w" -o /app/product product/rpc/./product.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/product /app/product
COPY --from=builder /app/etc /app/etc
COPY --from=builder /build/grpc_health_probe /bin/grpc_health_probe


EXPOSE 8080

CMD ["./product", "-f", "etc/product.yaml"]

