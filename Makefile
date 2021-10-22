BINARY_NAME=composite
build:
	go build -o ${BINARY_NAME} ./...

run:
	./${BINARY_NAME}
fmt:
	go fmt
clean:
	go clean
	rm ${BINARY_NAME}
all: fmt build run
