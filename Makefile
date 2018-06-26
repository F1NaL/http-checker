BINARY := http_cheker

.PHONY: linux
linux:
	mkdir -p build/{386,amd64}
	GOOS=linux GOARCH=amd64 go build -o build/amd64/$(BINARY) main.go
	GOOS=linux GOARCH=386 go build -o build/386/$(BINARY) main.go

.PHONY: darwin
darwin:
	mkdir -p build/{386,amd64}
	GOOS=darwin GOARCH=amd64 go build -o build/amd64/$(BINARY)_osx main.go

.PHONY: build
build:  linux darwin