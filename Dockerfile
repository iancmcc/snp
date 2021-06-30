FROM golang:1.16.0 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /dist/snp

FROM scratch
COPY --chown=0:0 --from=builder /dist /
USER 65534

ENTRYPOINT ["/snp"]
