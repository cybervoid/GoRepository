# {concierge,librarian}/Dockerfile

FROM golang:1.10

# In case of librarian, '/concierge' will be replaced with '/librarian'

ADD . /go/src/github.com/cybervoid/GoRepository/DocumentSearch/librarian

WORKDIR /go/src/github.com/cybervoid/GoRepository/DocumentSearch/librarian

RUN go install github.com/cybervoid/GoRepository/DocumentSearch/librarian

ENTRYPOINT /go/bin/librarian

EXPOSE 9000
