FROM golang:1.9.2-alpine

### START: Setting Environment ###

	ENV GOPATH /go
	ENV GOAPP template
	ENV GOENV kube
	ENV PATH $GOPATH/bin:$PATH
	ADD conf/bin $GOPATH/bin

	RUN apk add --no-cache \
        libc6-compat
	
### END: Setting Environment ###

### START: add source ###

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

	RUN go get github.com/beego/bee
	WORKDIR /go/src/template

### END: Initialize dependency ###

CMD ["/go/src/template/template"]