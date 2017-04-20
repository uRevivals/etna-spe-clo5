FROM golang

ADD . /go/src/reservation

RUN cd /go/src/reservation && go get && go install

ENTRYPOINT /go/bin/reservation