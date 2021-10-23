BINARY_NAME=$(shell basename $(CURDIR))
build:
	go build -o ${BINARY_NAME} ./...
tidy:
	go mod tidy
run:
	./${BINARY_NAME}
fmt:
	go fmt
test:
	go test
clean:
	go clean
	rm -rf ${BINARY_NAME}
all: fmt tidy build test run
