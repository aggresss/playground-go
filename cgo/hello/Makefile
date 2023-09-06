all:
	gcc -o libsum.so -shared sum.c
	ar -crv libsum.a sum.c
ifeq ($(shell uname -s), Darwin)
	go build
else
	go build -ldflags "-linkmode external -extldflags -static"
endif
