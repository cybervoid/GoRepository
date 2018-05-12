FROM golang:1.10
# the base image we want to use to build our docker image from
# specialized for golang
ADD . /go/src/hello
# copy files and folders from our system onto the docker image
RUN go install hello
# create executable binary for our project with the command
ENV NAME RyRy
# environment variable NAME will be picked up by the program 'hello' and printed to console.ENTRYPOINT /go/bin/hello
# command to execute when we start the container

# EXPOSE 9000
# generally used for network applications. allows us to connect to the application running inside the container from host system's localhost
ENTRYPOINT /go/bin/hello
