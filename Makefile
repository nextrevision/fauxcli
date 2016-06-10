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
	gox -os="darwin linux"

build-ci:
	go get github.com/mitchellh/gox
	sudo chown -R ${USER}: /usr/local/go
	$(MAKE) build
