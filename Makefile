TEST_RUN ?= com.go-crud/usecase/user com.go-crud/entity

.PHONY: all dep build clean

all: clean build

test: dep
	@echo "Testing..."
	@go test $(TEST_RUN) -v -coverprofile coverage.out

dep:
	@echo "Getting the dependencies"
	@go mod download

build: dep clean
	@echo "Building server"
	@go build -o ./cmd/api api/server.go

clean:
	@echo "Cleaning the last build"
	@rm -rf ./cmd
