FROM golang

ADD . /go/src/hotelObject

RUN cd /go/src/hotelObject && go get && go install

ENTRYPOINT /go/bin/hotelObject