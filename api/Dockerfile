FROM golang:1.10

ARG port

ADD . /go/src/github.com/JILeXanDR/docker-compose-golang-postgre-govendor/api

WORKDIR /go/src/github.com/JILeXanDR/docker-compose-golang-postgre-govendor/api

# install package manager Glide https://github.com/Masterminds/glide
RUN curl https://glide.sh/get | sh
RUN go get github.com/pilu/fresh

# install dependencies
CMD glide install && fresh -c runner.conf main.go

# build binary file
#RUN go install github.com/JILeXanDR/docker-compose-golang-postgre-govendor/api
# run binary when container starts
#ENTRYPOINT /go/bin/api

EXPOSE ${port}
