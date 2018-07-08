BUILD_BIN_NAME=brasilio
BUILD_DIR=build

COMMIT_ID := $(shell git rev-parse HEAD)

build:
	# Install build dependencies
	go get github.com/inconshreveable/mousetrap
	rm -Rf $(BUILD_DIR)
	# Linux x86
	env GOOS=linux 	 GOARCH=386 go build -o $(BUILD_DIR)/linux/$(BUILD_BIN_NAME) main.go
	# MacOS X x86
	env GOOS=darwin  GOARCH=386 go build -o $(BUILD_DIR)/macosx/$(BUILD_BIN_NAME) main.go
	# Windows x86
	env GOOS=windows GOARCH=386 go build -o $(BUILD_DIR)/windows/$(BUILD_BIN_NAME).exe main.go

deploy:
	# Compress build
	zip -r $(BUILD_DIR)/build.zip $(BUILD_DIR)/linux/ $(BUILD_DIR)/macosx/ $(BUILD_DIR)/windows/
	# Upload
	curl -F "commit_id=$(COMMIT_ID)" -F "upload_token=$(UPLOAD_TOKEN)" -F "builds.zip=@$(BUILD_DIR)/build.zip" $(UPLOAD_URL)