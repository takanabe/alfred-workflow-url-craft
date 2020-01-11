GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=url-craft

.PHONY: all
all: build

#
# Build binary
#
.PHONY: build
build:
		$(GOBUILD) -o workflow/$(BINARY_NAME) -v
