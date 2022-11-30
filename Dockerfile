FROM golang
MAINTAINER yawntee@qq.com

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV WF_DIR=/data

VOLUME /data

COPY . /root/wf_api
WORKDIR /root/wf_api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o /usr/local/bin/wf_backend wf_api/server

WORKDIR /root/wf_api/server

EXPOSE 8888

ENTRYPOINT ["wf_backend"]

