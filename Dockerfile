FROM ysitd/glide as builder

ADD . /go/src/github.com/ysitd-cloud/app-controller

WORKDIR /go/src/github.com/ysitd-cloud/app-controller

RUN glide install -v --force && \
    make all

FROM alpine:3.5

COPY --from=builder /go/src/github.com/ysitd-cloud/app-controller/controller /

ENTRYPOINT ["/controller"]