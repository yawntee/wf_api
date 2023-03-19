# Compile stage
FROM golang as build-env
MAINTAINER yawntee@qq.com

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

COPY . /dockerdev
WORKDIR /dockerdev

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS="-buildvcs=false" go build -ldflags '-w -s' -o /server

# Final stage

FROM debian:buster

COPY --from=build-env /server /wf
WORKDIR /wf

VOLUME /wf/data

EXPOSE 8888

CMD ["/server"]