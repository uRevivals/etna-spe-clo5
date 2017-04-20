FROM golang

ADD . /go/src/hotelObject

RUN cd hotelObject && go get /go/src/hotelObject && go install

ENTRYPOINT /go/bin/hotelObject

EXPOSE 1234