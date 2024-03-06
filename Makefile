build:
	@go build -o bin/quickserver

run: build
	./bin/quickserver

test:
	go test -v ./..