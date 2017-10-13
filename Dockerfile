FROM scratch

COPY dist /

ENTRYPOINT ["/manager"]