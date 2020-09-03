## Prerequisites
Install these first:
- Golang 1.14: https://golang.org/
- IBM db driver for Golang: https://github.com/ibmdb/go_ibm_db

> You must download this repository to your $GOPATH

You can check $GOPATH with command:

```go env GOPATH```


## Start program
Download Go libraries by this command:

```go get -v```

```go mod init```

Then run

```go run server.go util.go model.go constant.go controller.go dropdown_repository.go jwt.go```

or

```go build -o app .```

If you're using Windows then run:

```app.exe```


## You can build Docker container instead

- cd to project path then

```docker build --tag rest-api-with-golang-gin:1.0.0_x .```

- start the container

```docker run --name rest-api-with-golang-gin -p 8080:8080 --detach rest-api-with-golang-gin:1.0.0_x```