FROM golang:alpine

RUN apk add --update git

ARG WORKDIR=/go/src/github.com/hypertrace/agent-config

RUN mkdir -p $WORKDIR/tools/env-vars-generator

WORKDIR $WORKDIR

COPY . $WORKDIR/tools/env-vars-generator

# Need the go modules in the WORKDIR
COPY go.* .

ENTRYPOINT ["/usr/local/go/bin/go", "run", "tools/env-vars-generator/main.go"]