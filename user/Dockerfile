FROM golang

ADD . /go/src/user

RUN cd /go/src/user && go get && go install

ENTRYPOINT /go/bin/user