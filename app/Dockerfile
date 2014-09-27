FROM google/golang

WORKDIR /gopath/src/app
ADD . /gopath/src/app
RUN go get app

CMD ["/gopath/bin/app"]