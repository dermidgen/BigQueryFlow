FROM golang:1.14-alpine
ENV CGO_ENABLED=0
ENV GO111MODULE=on

ARG BQFLOW_VERSION=unknown

WORKDIR /go/src/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -v -o /usr/local/bin/bqflow -ldflags="-s -w -X github.com/dermidgen/BigQueryFlow/run.version=$BQFLOW_VERSION" ./

ENTRYPOINT ["bqflow"]
