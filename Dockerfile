FROM alpine:3.8

LABEL maintainer="wagnst"

ARG BUILD_VERSION

RUN apk update
RUN apk upgrade
# needed to run a go binary inside alpine
RUN apk add --no-cache \
        libc6-compat

RUN echo "Running Dockerbuild for version $BUILD_VERSION"
ADD ./bin/go-docker-sample-$BUILD_VERSION /
RUN ln -s /go-docker-sample-$BUILD_VERSION /go-docker-sample
RUN chmod +x /go-docker-sample-$BUILD_VERSION
RUN chmod +x /go-docker-sample
RUN ls -ltra /

ENTRYPOINT ["/go-docker-sample"]

EXPOSE 8000