BINARY='gifanimecreator.wasm'

all: build

build:
	GOOS=js GOARCH=wasm go build -o ${BINARY} ./cmd/gifanimecreator

build-prod:
	GOOS=js GOARCH=wasm go build -o ${BINARY} -ldflags "-s -w" ./cmd/gifanimecreator
