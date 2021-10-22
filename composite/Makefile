BINARY_NAME=composite_bin
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
	rm ${BINARY_NAME}
all: fmt tidy build run
