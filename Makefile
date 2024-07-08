.PHONY: test build

test: build
	@build/generate-docs -filepath test/open-api-spec.yaml -output build/spec

build:
	@go build -o build/generate-docs  main.go
