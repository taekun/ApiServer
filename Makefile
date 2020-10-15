.PHONY: build
build:
		go build -v ./cmd/apiserver

.DEFAULT_GDAL != build