BUILD_VERSION=$(shell git describe --always --long --abbrev=7)
all:
	go build -o main.out -tags debug -ldflags '-X main.Version=${BUILD_VERSION}' main.go
