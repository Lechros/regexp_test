# rure-go
FROM rust AS rust-builder

WORKDIR /app
RUN git clone --depth 1 https://github.com/rust-lang/regex
RUN cargo build --release --manifest-path ./regex/regex-capi/Cargo.toml

FROM golang:1.23 AS go-builder

# go-pcre, gohs
RUN apt -y update \
    && apt install -y --no-install-recommends pkg-config libpcre3-dev

# rure-go
ENV CGO_LDFLAGS="-L/app/regex/target/release"
ENV LD_LIBRARY_PATH="/app/regex/target/release"
WORKDIR /app
COPY --from=rust-builder /app/regex/target/release /app/regex/target/release

# go build
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY *.json ./
COPY main.go main_test.go ./
RUN go build -o app ./...

ENTRYPOINT ["go", "test", "-bench", "."]

#FROM golang:1.23
#
#WORKDIR /app
## rure-go
#COPY --from=go-builder /app/regex/target/release/*.so /usr/lib/
#
#COPY --from=go-builder /go/src/app/app ./
#COPY *.json ./
#
#ENTRYPOINT ["./app"]