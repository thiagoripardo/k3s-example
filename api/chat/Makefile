.PHONY: build clean deploy test

GO_BUILD := go build -ldflags '-s -w'

clean:
	rm -rf ./bin

build: clean
	export GO111MODULE=on
	export CGO_ENABLED=0
	
	$(GO_BUILD) -o bin/app main.go

test:
	go test -v -failfast ./...
