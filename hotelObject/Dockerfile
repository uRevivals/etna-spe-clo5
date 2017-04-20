FROM golang

ADD . /go/src/github.com/uRevivals/microservice/hotelObject

RUN go install github.com/uRevivals/microservice/hotelObject

ENTRYPOINT /go/bin/hotelObject

EXPOSE 1234