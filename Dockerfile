FROM ysitd/glide

WORKDIR /go/src/github.com/ysitd-cloud/app-controller

ADD . /go/src/github.com/ysitd-cloud/app-controller

RUN glide install -v --skip-test && \
    go install

CMD ["app-controller"]
