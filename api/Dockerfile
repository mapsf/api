FROM golang:1.10

ARG DOCKER_GOLANG_PORT
ARG REPOSITORY_NAME

ADD . /go/src/github.com/${REPOSITORY_NAME}

WORKDIR /go/src/github.com/${REPOSITORY_NAME}

# install package manager Glide https://github.com/Masterminds/glide
RUN curl https://glide.sh/get | sh
RUN go get github.com/pilu/fresh

# install dependencies
CMD glide up && fresh -c runner.conf main.go

# build binary file
#RUN go install github.com/${REPOSITORY_NAME}
# run binary when container starts
#ENTRYPOINT /go/bin/api

EXPOSE ${DOCKER_GOLANG_PORT}
