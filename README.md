# covid-summary
simple JSON API to summarize COVID-19 stats using Golang with Gin framework

## Directory structure
```
├── README.md
├── api
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── middleware
├── mockdata
├── model
├── router
└── service
```

## Getting started

### Install Golang
This project use Go 1.16 to develop, make sure you have Go 1.16 or higher installed.

https://golang.org/doc/install

### How to run
```
go run cmd/main.go
```

### Testing
In current project I focus testing in core function and each test file must coverage more than 80%

From the project root, run:
```
go test -v -cover ./...
```
