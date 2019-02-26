FROM alpine:3.9 AS builder

RUN apk --no-cache update && \

    apk --no-cache upgrade && \

    apk --no-cache add \
    go git ca-certificates musl-dev && \

    adduser -D -g '' docker && \
    mkdir -pv /docker && \
    chown -R docker:docker /docker

USER docker
ENV GOPATH=/docker/gopath \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . /docker
WORKDIR /docker

RUN go get -v -u \
    github.com/thoj/go-ircevent && \
    go build -a -installsuffix cgo -ldflags \
    '-extldflags "-static"' -o /docker/nixy .

FROM scratch

ENV GOPATH=/docker/gopath

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /tmp /tmp
COPY --from=builder /etc/ssl/certs/ca-certificates.crt \
    /etc/ssl/certs/

COPY --from=builder /docker/gopath \
    /docker/gopath
COPY --from=builder /docker/nixy \
    /docker/nixy

USER docker

CMD ["-help"]
ENTRYPOINT ["/docker/nixy"]
