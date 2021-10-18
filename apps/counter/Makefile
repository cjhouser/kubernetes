SHELL = /bin/sh

# Detect number of CPUs available for parallel jobs.
NPROCS := 1
OS := $(shell uname -s)
ifeq ($(OS),Linux)
NPROCS := $(shell grep -c ^processor /proc/cpuinfo)
endif
ifeq ($(OS),Darwin)
NPROCS := $(shell sysctl hw.ncpu | grep -o '[0-9]\+')
endif
MAKEFLAGS += -j$(NPROCS)

OBJECTS = counter.go

build:
	go build -o build/counter $(OBJECTS)

run:
	go run $(OBJECTS)

compile: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64
	
linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o build/counter-linux-amd64 $(OBJECTS)

linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o build/counter-linux-arm64 $(OBJECTS)

darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o build/counter-darwin-amd64 $(OBJECTS)

darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o build/counter-darwin-arm64 $(OBJECTS)
