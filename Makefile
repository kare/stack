
GENERATED_FILES = int.go string.go complex128.go float32.go int16.go int64.go uint16.go uint64.go byte.go complex64.go float64.go int32.go int8.go rune.go uint.go uint32.go uint8.go

.PHONY: generate gends test build vet

test: build
	@go test

build: generate
	@go build

vet: test
	gofmt -s -w $(GENERATED_FILES) stack.go
	golint $(GENERATED_FILES) stack.go
	govet $(GENERATED_FILES) stack.go
	errcheck $(GENERATED_FILES) stack.go
	staticcheck $(GENERATED_FILES) stack.go
	unused $(GENERATED_FILES) stack.go
	gosimple $(GENERATED_FILES) stack.go

gends:
	@go install kkn.fi/cmd/gends

generate: gends
	@go generate

clean:
	@rm -rf $(GENERATED_FILES)
