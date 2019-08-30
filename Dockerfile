FROM golang:1.12

# RUN mkdir /app

RUN go get -u github.com/golang/dep/cmd/dep

COPY . /go/src/app

WORKDIR  /go/src/app/cmd/email

RUN dep ensure 
RUN go build -o main .

EXPOSE 8080

CMD ["/go/src/app/cmd/email/main"]
