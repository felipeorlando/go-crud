GREEN=\033[0;32m

.PHONY: all dep build clean

all: clean build

dep:
	@echo "Getting the dependencies"
	@go mod download

build: dep clean
	@echo "Building server"
	@go build -o ./cmd/api api/main.go

clean:
	@echo "Cleaning the last build"
	@rm -rf ./cmd
