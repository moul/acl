FROM    golang:1.7.3
COPY    . /go/src/github.com/moul/acl
WORKDIR /go/src/github.com/moul/acl
CMD     ["acl"]
EXPOSE  8000 9000
RUN     make install
