# {concierge,librarian}/Dockerfile

FROM golang:1.10

# In case of librarian, '/concierge' will be replaced with '/librarian'

ADD . /go/src/github.com/cybervoid/GoRepository/DocumentSearch/concierge

WORKDIR /go/src/github.com/cybervoid/GoRepository/DocumentSearch/concierge

RUN go install github.com/cybervoid/GoRepository/DocumentSearch/concierge

ENTRYPOINT /go/bin/concierge

EXPOSE 9000
