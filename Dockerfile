FROM golang:1.12.0

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o ButteAir

CMD ["/app/ButteAir", "-prod"]