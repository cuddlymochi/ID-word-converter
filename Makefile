export GOPATH := $(shell pwd)
default: build

init:
	rm -f bin/server bin/main bin/T9-server
	@cd src/main && go get

build: init
	go build -o bin/T9-server src/main/main.go 

run: build
	@pkill ^T9-server$ || :
	bin/T9-server>log.txt 2>&1 &

log: run
	tail -f -n2 log.txt
