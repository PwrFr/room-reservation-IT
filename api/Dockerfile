FROM golang

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

EXPOSE 3001
CMD ["/app/main"]