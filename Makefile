# Focus — Build System
# Requires: Go 1.21+, go-winres (for manifest embedding)
# Install go-winres: go install github.com/tc-hib/go-winres@latest

APP      = focus
OUTPUT   = dist/$(APP).exe
MAIN     = ./cmd/focus

.PHONY: all build deps clean icon

all: deps build

deps:
	go mod tidy
	go mod download

# Embed manifest + icon into a syso file (picked up automatically by go build)
icon:
	@if exist assets\focus.ico ( \
		go-winres make --in focus.rc --out cmd/focus/rsrc.syso \
	) else ( \
		echo "Warning: assets/focus.ico not found, skipping resource embedding" \
	)

build: icon
	@if not exist dist mkdir dist
	GOOS=windows GOARCH=amd64 go build \
		-ldflags="-H windowsgui -s -w" \
		-o $(OUTPUT) \
		$(MAIN)
	@echo Built: $(OUTPUT)

# Build without hiding console (useful for debugging)
build-debug:
	@if not exist dist mkdir dist
	GOOS=windows GOARCH=amd64 go build \
		-o dist/$(APP)-debug.exe \
		$(MAIN)

clean:
	rm -rf dist/
	rm -f cmd/focus/rsrc.syso
