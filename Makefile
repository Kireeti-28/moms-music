BINARY_NAME=blog
OS=windows

build:
	GOARCH=amd64 GOOS=${OS} go build -o ./bin/${BINARY_NAME}

run: build
	./bin/${BINARY_NAME}

clean:
	go clean
	rm ./bin/${BINARY_NAME}

test:
	go test ./...