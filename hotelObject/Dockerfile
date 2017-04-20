FROM golang

ADD . /go/src/hotelObject

RUN cd /go/src/hotelObject && go get /go/src/hotelObject && go install

ENTRYPOINT /go/bin/hotelObject

EXPOSE 1234