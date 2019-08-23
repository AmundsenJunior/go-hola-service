# go-hola-service

[![Build Status](https://travis-ci.org/AmundsenJunior/go-hola-service.svg?branch=master)](https://travis-ci.org/AmundsenJunior/go-hola-service)

*Go web server as an example hello world service for writing applications in Go.*

## Structure

### Application

* `main.go` starts, runs, and stops the whole Go application
* `app.go` holds the App type of `mux.Router` and its initializers of routes/handlers
* `handlers.go` defines the handler functions of the service

### Other

* `main_test.go` defines the set of service tests
* `Dockerfile` executes a multistage build of the Go binary and the app container

## Development 

### Build and run as a Docker container

Using the multistage build pattern in Docker, this project's `Dockerfile` pulls module dependencies, executes `go test`,
and `go build` to create the application binary. This binary only is copied over to an `alpine` image and run as a
service with its port exposed.

```shell script
$ docker build -t go-hola-svc .
$ docker run -d -p 8000:8000 --name go-hola-svc go-hola-svc
$ docker logs -f go-hola-svc
$ curl localhost:8000/hello
```
