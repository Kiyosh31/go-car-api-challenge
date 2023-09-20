OS_NAME=$(shell uname -s | tr A-Z a-z) # linux|mac|windows
BINARY_NAME_LINUX=cars_api-linux
BINARY_NAME_WINDOWS=cars_api-windows
BINARY_NAME_MAC=cars_api-darwin

build:
	GOARCH=amd64 GOOS=linux go build -o bin/$(BINARY_NAME_LINUX) main.go
	GOARCH=amd64 GOOS=darwin go build -o bin/$(BINARY_NAME_MAC) main.go
	GOARCH=amd64 GOOS=windows go build -o bin/$(BINARY_NAME_WINDOWS) main.go

dev:
	go run main.go

run: build
		@if [ $(OS_NAME) = "linux" ]; then\
			./bin/$(BINARY_NAME_LINUX);\
		fi
		@if [$(OS_NAME) = "windows" ]; then\
			./bin/$(BINARY_NAME_WINDOWS);\
		fi
		@if [$(OS_NAME) = "darwin" ]; then\
			./bin/$(BINARY_NAME_MAC);\
		fi

clean:
	go clean
	rm -rf bin/

test:
	go test -v ./...

test_coverage:
	go test -v ./... -coverprofile=coverage.out

vet:
	go vet

mod:
	go mod download