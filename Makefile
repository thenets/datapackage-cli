BUILD_BIN_NAME=brasilio
BUILD_DIR=builds

build:
	# Install build dependencies
	go get github.com/inconshreveable/mousetrap
	mkdir -p $(BUILD_DIR)
	# Linux x86
	env GOOS=linux 	 GOARCH=386 go build -o $(BUILD_DIR)/linux/$(BUILD_BIN_NAME) main.go
	# MacOS X x86
	env GOOS=darwin  GOARCH=386 go build -o $(BUILD_DIR)/macosx/$(BUILD_BIN_NAME) main.go
	# Windows x86
	env GOOS=windows GOARCH=386 go build -o $(BUILD_DIR)/windows/$(BUILD_BIN_NAME).exe main.go