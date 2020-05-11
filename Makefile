.PHONY: build clean tool lint help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=juice-server

# Rules
all: test build
build:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -v
	mkdir ./build/scripts
	cp ./scripts/mp3downloader.py ./build/scripts/mp3downloader.py