FROM alpine:3.5

COPY controller /

ENTRYPOINT ["/controller"]