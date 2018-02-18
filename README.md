# gRPC Go Sandbox

The sandbox of the gRPC apps implemented in Go.

## Requirements

* Go
* [dep](https://github.com/golang/dep)
* [proto-easy](https://github.com/peter-edge/protoeasy-go)

## How to setup

### Server app

```
$ cd server
$ dep ensure
$ go generate
```

### Client app

```
$ cd client
$ dep ensure
```

### Gateway app

```
$ cd gateway
$ dep ensure
```

## Usage

Launch a gRPC server.

```
$ go run server/main.go
```

Execute a gRPC client that call the Greeter service of gRPC server.

```
$ go run client/main.go
2018/02/18 10:27:23 Greeting: Hello world
```

Launch a [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

```
$ go run gateway/main.go
```

And request `GET http://localhost:3000/v1/greeter`.

```
$ curl http://localhost:3000/v1/greeter
{"message":"Hello "}
```
