# Build the Go API
FROM    golang:1.14.4-stretch AS go_builder
RUN     apt-get update && apt-get install -y --no-install-recommends libxml2
COPY    . /go/src
WORKDIR /go/src
RUN     git config --global http.sslVerify false
RUN     go get -d -v -insecure .
WORKDIR /go/src/github.com/ibmdb/go_ibm_db/installer
RUN     go run setup.go
ENV     DB2HOME=/go/src/github.com/ibmdb/go_ibm_db/installer/clidriver
ENV     CGO_CFLAGS=-I$DB2HOME/include
ENV     CGO_LDFLAGS=-L$DB2HOME/lib
ENV     LD_LIBRARY_PATH=$DB2HOME/lib
WORKDIR /go/src
RUN     go build -o app .
EXPOSE  8080
CMD     ["./app"]