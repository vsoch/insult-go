
all:
	gofmt -w *.go
	GOOS=js GOARCH=wasm go build -o docs/main.wasm
