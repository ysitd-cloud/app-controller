FROM scratch

COPY controller /

ENTRYPOINT ["/controller"]