FROM golang:1.11
LABEL maintainer="Nirmoy Das <nirmoy.aiemd@gmail.com>"
WORKDIR $GOPATH/src/github.com/nirmoy/consul-operator
COPY . .
RUN GO111MODULE=on go get -d -v ./...
RUN GO111MODULE=on go install -v ./...
CMD ["consul-operator"]
