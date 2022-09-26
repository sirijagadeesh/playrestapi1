.PHONY = default_build build version clean lint test build run

default_build: build

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

version:	## version of go installed in mechine.
version:
	@go version

clean:		## verifying formatting, race and dependency verification.
clean: version
	@go fmt ./...				
	@gofumpt -l -w . 
	@go vet ./...
	@go mod tidy
	@go mod verify

lint:		## linting the code.
lint: clean
	@golangci-lint run --enable-all --tests=false ./...

test:		## run test case if any.
test: lint
	@go test -v ./...

build:		## build the binay in ./bin/
build: test
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/

run:		## run the binary crated.
run: build
	@echo "--------- running code ---------"
	@time ./bin/playrestapi1

image:		## build docker images with name:tag playrestapi1:latest
image:
	@docker build -f ./docker/dockerfile -t playrestapi1:latest .
	@echo "image build completed"