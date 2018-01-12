FROM ysitd/glide as builder

ADD . /go/src/code.ysitd.cloud/component/deployer

WORKDIR /go/src/code.ysitd.cloud/component/deployer

RUN apk add --no-cache make && \
    glide install -v --force && \
    make all

FROM alpine:3.5

COPY --from=builder /go/src/code.ysitd.cloud/component/deployer/controller /

ENTRYPOINT ["/controller"]