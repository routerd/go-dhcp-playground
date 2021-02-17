FROM golang:1.15 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY cmd cmd

RUN CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    go build -o go-dhcp-playground ./cmd/playground

FROM nicolaka/netshoot:latest
WORKDIR /
COPY --from=builder /workspace/go-dhcp-playground go-dhcp-playground

ENTRYPOINT ["/go-dhcp-playground"]
CMD ["eth0", "0.0.0.0"]