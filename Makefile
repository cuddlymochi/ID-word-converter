export GOPATH := $(shell pwd)
default: build

init:
	rm -f bin/server bin/main bin/Convert
	@cd src/main && go get

build: init
	go build -o bin/Convert src/main/main.go 

run: build
	@pkill ^Convert$ || :
	bin/Convert>log.txt 2>&1 &

log: run
	tail -f -n2 log.txt
