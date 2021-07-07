BUILD_DIR=./build

clean:
	@rm -r ${BUILD_DIR}

fmt:
	@gofmt -s -w .

lint:
	@golangci-lint run .

test:
	go test .

check: fmt lint test

${BUILD_DIR}:
	@mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/motoblog-pipeline main.go
