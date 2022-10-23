FROM golang:1.19.2 

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o cmd/server . 
CMD ["/app/server"]