FROM golang:1.9.2

### START: Setting Environment ###

	ENV GOPATH /go
	ENV PATH $GOPATH/bin:$PATH
	
### END: Setting Environment ###

### START: Set Date Time ###

	RUN ln -sf /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
	RUN echo "Asia/Bangkok" > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata
	ENV TZ=Asia/Bangkok
	RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

### END Set Date Time ###

### START: add source ###

	RUN mkdir -p /go/src/template
	RUN mkdir -p /go/src/template/logs

	ADD . /go/src/template

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