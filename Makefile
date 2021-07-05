BUILD_DIR=./build

clean:
	@rm -r ${BUILD_DIR}

${BUILD_DIR}:
	@mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/motoblog-pipeline main.go
