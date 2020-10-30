# This Dockerfile serves solely for the purpose of being able to run the golang tasks without
# the need of having Go installed locally.

FROM golang

WORKDIR /go/src/github.com/hypertrace/agent-config/go

COPY config.proto ../config.proto

RUN go get -u github.com/iancoleman/strcase
RUN go get -u github.com/tallstoat/pbparser

ENTRYPOINT ["go"]