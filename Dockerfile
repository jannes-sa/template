FROM golang:1.9.2-alpine

### START: Setting Environment ###

	ENV GOPATH /go
	ENV PATH $GOPATH/bin:$PATH
	
### END: Setting Environment ###

### START: add source ###

	RUN mkdir -p /go/src/template
	RUN mkdir -p /go/src/template/logs

	ADD conf /go/src/template
	ADD storages /go/src/template
	ADD database /go/src/template

### END: add source ###


### START: Initialize dependency ###

	RUN go get github.com/beego/bee
	WORKDIR /go/src/template

### END: Initialize dependency ###

### START: Build Package ###
	RUN go build
	# RUN /go/src/template/runtime

### END: Build Package ###

CMD ["/go/src/template/template"]