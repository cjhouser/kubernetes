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
	docker buildx build --push --platform linux/arm64,linux/amd64 -t cjhouser/counter:latest -t cjhouser/counter:1.0.0 .
