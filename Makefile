MAJOR_VERSION = 1
MINOR_VERSION = 2
GIT_VERSION = 1.2.0
GIT_COMMIT = $(shell git log -1 --pretty=format:"%H")
BUILD_DATE = $(shell date +%Y-%m-%dT%H:%M:%S%z)
BUILD_VAR = -X version.gitCommit=$(GIT_COMMIT) -X version.buildDate=$(BUILD_DATE) -X version.major=$(MAJOR_VERSION) -X version.minor=$(MINOR_VERSION) -X version.gitVersion=$(GIT_VERSION)
BUILD_FLAG = -ldflags "$(BUILD_VAR)"

all: controller

.PHONY: clean

controller:
	go build $(BUILD_FLAG) -o controller