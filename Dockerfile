FROM scratch

COPY manager /

ENTRYPOINT ["/app-controller"]