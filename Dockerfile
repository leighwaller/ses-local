FROM golang:alpine

ENV INSTALL /go/src/ses-local
WORKDIR ${INSTALL}

ADD . $INSTALL

RUN go install .

ENTRYPOINT /go/bin/ses-local

EXPOSE 8080
