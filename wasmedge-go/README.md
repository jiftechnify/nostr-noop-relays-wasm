# Nostr no-op relay for WasmEdge (in Go)

## How to Build and Run
Prerequisite: Go >= 1.21 and `make` (build tool) have been installed.

1. Follow the steps [here](https://wasmedge.org/docs/start/install) to install WasmEdge.
2. Build and run the Wasm module by following commands: 

```bash
make build
# it executes: GOARCH=wasm GOOS=wasip1 go build -o noop-nrelay-wasmedge-go.wasm main.go

wasmedge noop-nrelay-wasmedge-go.wasm
```
