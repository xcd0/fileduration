BIN           := `grep "module" ./go.mod | sed 's/module //'`
REVISION      := `git rev-parse --short HEAD`
FLAG          :=  -a -tags netgo -trimpath -ldflags='-s -w -extldflags="-static" -buildid='
RESOURCE_DIR  := resources
BINDATA_FILE  := bindata.go
SOURCE_FILES  := go.mod go.sum *.go makefile .gitignore readme.md

.PHONY: all build clean release update-binary

all:
	cat ./makefile | grep '^[^ ]*:$$'
build:
	go build -o $(BIN).exe
release:
	make build
	GOOS=windows go build $(FLAG) -o $(BIN).exe
	upx --lzma $(BIN).exe
	echo Success!

