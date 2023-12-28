.PHONY: build
build: noop-nrelay-wasmedge-go.wasm

noop-nrelay-wasmedge-go.wasm: main.go nostr/nostr.go
	GOARCH=wasm GOOS=wasip1 go build -o noop-nrelay-wasmedge-go.wasm main.go

.PHONY: clean
clean:
	rm -f noop-nrelay-wasmedge-go.wasm
