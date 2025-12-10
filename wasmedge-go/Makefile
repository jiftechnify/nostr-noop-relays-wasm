noop-nrelay-wasmedge-go.wasm: main.go nostr/nostr.go go.mod
	GOARCH=wasm GOOS=wasip1 go build -o noop-nrelay-wasmedge-go.wasm main.go

.PHONY: build
build: noop-nrelay-wasmedge-go.wasm

.PHONY: run
run: noop-nrelay-wasmedge-go.wasm
	wasmedge noop-nrelay-wasmedge-go.wasm

.PHONY: clean
clean:
	rm -f noop-nrelay-wasmedge-go.wasm
