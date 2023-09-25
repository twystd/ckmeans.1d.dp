.PHONY: integration-tests

all: test      \
	 benchmark \
     coverage

clean:
	go clean

format: 
	go fmt ./...

build: format
	go build -o bin ./...

debug: build
	go test ./...

test: build
	go test ./...

run: build
	./bin/ckmeans --debug ../test/taps.txt

integration-tests: build
	mkdir -p ../test/tmp
	rm -f ../test/tmp/*
	./bin/ckmeans --out ../test/tmp/taps.txt ../test/taps.txt
	diff ../test/expected ../test/tmp

vet: build
	go vet ./...

lint: build
	golint ./...

benchmark: build
	go test -bench ./...

coverage: build
	go test -cover ./...


