TEST?=

default: test

test:
	go test -v $(TEST) $(TESTARGS)
	go vet -v

cover:
	go test $(TEST) -covermode=count -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

build:
	gox -os="darwin linux" -osarch="darwin/386 darwin/amd64 linux/386 linux/amd64"

build-ci:
	go get github.com/mitchellh/gox
	sudo chown -R ${USER}: /usr/local/go
	$(MAKE) build
