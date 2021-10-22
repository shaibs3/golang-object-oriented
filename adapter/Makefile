BINARY_NAME=adapter_bin
build:
	go build -o ${BINARY_NAME} ./...
tidy:
	go mod tidy
run:
	./${BINARY_NAME}
fmt:
	go fmt
clean:
	go clean
	rm -rf ${BINARY_NAME}
all: fmt tidy build run
