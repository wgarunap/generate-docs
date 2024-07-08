.PHONY: test build

test: build
	@build/bin/generate-docs -filepath test/open-api-spec.yaml -output build/spec

build:
	@go build -o build/bin/generate-docs  main.go
