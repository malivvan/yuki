.PHONY: build
.DEFAULT_GOAL := build

BUILD = go build -trimpath -ldflags "-s -w -buildid=" -tags=netgo,osusergo -o build/yuki

clean:
	mkdir -p build
	rm -f build/*

test:
	go test ./modules/buffer
	go test ./modules/console
	go test ./modules/eventloop
	go test ./modules/library
	go test ./modules/process
	go test ./modules/require
	go test ./modules/url
	go test ./modules/util


linux_amd64:
	 CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(BUILD)_linux_amd64 ./cmd

linux_arm64:
	 CGO_ENABLED=1 GOOS=linux GOARCH=arm64 $(BUILD)_linux_arm64 ./cmd

freebsd_amd64:
	 CGO_ENABLED=1 GOOS=freebsd GOARCH=amd64 $(BUILD)_freebsd_amd64 ./cmd

freebsd_arm64:
	 CGO_ENABLED=1 GOOS=freebsd GOARCH=arm64 $(BUILD)_freebsd_arm64 ./cmd

openbsd_amd64:
	CGO_ENABLED=1 GOOS=openbsd GOARCH=amd64 $(BUILD)_openbsd_amd64 ./cmd

openbsd_arm64:
	CGO_ENABLED=1 GOOS=openbsd GOARCH=arm64 $(BUILD)_openbsd_arm64 ./cmd

netbsd_amd64:
	CGO_ENABLED=1 GOOS=netbsd GOARCH=amd64 $(BUILD)_netbsd_amd64 ./cmd

netbsd_arm64:
	CGO_ENABLED=1 GOOS=netbsd GOARCH=arm64 $(BUILD)_netbsd_arm64 ./cmd

dragonfly_amd64:
	CGO_ENABLED=1 GOOS=dragonfly GOARCH=amd64 $(BUILD)_dragonfly_amd64 ./cmd

darwin_amd64:
	 CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(BUILD)_darwin_amd64 ./cmd

darwin_arm64:
	 CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 $(BUILD)_darwin_arm64 ./cmd

windows_amd64:
	 CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(BUILD)_windows_amd64.exe ./cmd

windows_arm64:
	 CGO_ENABLED=1 GOOS=windows GOARCH=arm64 $(BUILD)_windows_arm64.exe ./cmd

js_wasm:
	CGO_ENABLED=0 GOOS=js GOARCH=wasm $(BUILD)_js_wasm ./cmd

build: clean linux_amd64 linux_arm64 freebsd_amd64 freebsd_arm64 darwin_amd64 darwin_arm64 windows_amd64 windows_arm64 openbsd_amd64 openbsd_arm64 netbsd_amd64 netbsd_arm64 dragonfly_amd64 js_wasm

info:
	file build/*
	du -h build/*
