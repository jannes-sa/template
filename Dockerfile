FROM golang:1.9.2-alpine

### START: Setting Environment ###

	ENV GOPATH /go
	ENV GOAPP template
	ENV GOENV kube
	ENV PATH $GOPATH/bin:$PATH
	ADD conf/bin $GOPATH/bin

	RUN apk add --no-cache \
        libc6-compat
	RUN apk add --no-cache tzdata
### END: Setting Environment ###

### START: add source ###
	RUN mkdir -p /go/src/template/helper
	ADD helper /go/src/template/helper

	RUN mkdir -p /go/src/template/vendor/github.com/astaxie
	ADD vendor/github.com/astaxie /go/src/template/vendor/github.com/astaxie

	RUN mkdir -p /go/src/template/vendor/github.com/go-sql-driver/mysql
	ADD vendor/github.com/go-sql-driver/mysql /go/src/template/vendor/github.com/go-sql-driver/mysql

	RUN mkdir -p /go/src/template/vendor/github.com/lib/pq
	ADD vendor/github.com/lib/pq /go/src/template/vendor/github.com/lib/pq

	RUN mkdir -p /go/src/template
	RUN mkdir -p /go/src/template/logs

	RUN mkdir -p /go/src/template/conf	
	ADD conf /go/src/template/conf

	RUN mkdir -p /go/src/template/storages
	ADD storages /go/src/template/storages
	
	RUN mkdir -p /go/src/template/database
	ADD database /go/src/template/database
	
	ADD template /go/src/template

	
### END: add source ###


### START: Initialize dependency ###

	# RUN apk add --no-cache git mercurial \
    # && go get github.com/beego/bee \
    # && apk del git mercurial
	WORKDIR /go/src/template
	RUN ls
### END: Initialize dependency ###

CMD ["/go/src/template/template"]