BUILD_VAR :=

GIT_COMMIT = $(shell git log -1 --pretty=format:"%H")
BUILD_VAR += -X version.gitCommit=$(GIT_COMMIT)

BUILD_DATE = $(shell date +%Y-%m-%dT%H:%M:%S%z)
BUILD_VAR += -X version.buildDate=$(BUILD_DATE)
MAJOR_VERSION = 1
MINOR_VERSION = 2
BUILD_VAR += -X version.major=$(MAJOR_VERSION)
BUILD_VAR += -X version.minor=$(MINOR_VERSION)

GIT_VERSION = 1.2.0
BUILD_VAR += -X version.gitVersion=$(GIT_VERSION)

GIT_TREE_STATE = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
BUILD_VAR += -X version.gitTreeState=$(GIT_TREE_STATE)

BUILD_FLAG = -ldflags "$(BUILD_VAR)"

all: controller

.PHONY: clean

controller:
	go build $(BUILD_FLAG) -o controller