FROM golang:1.13 as build

WORKDIR /home/lincx_py
ADD . .
RUN go build .
EXPOSE 8080
ENTRYPOINT ["./webproject"]
