# Nostr no-op relay for WasmEdge (in Rust)

## How to Build and Run
1. Follow the steps [here](https://wasmedge.org/docs/develop/rust/setup/) to setup WasmEdge & Rust toolchain.
2. Build and run the Wasm module by following commands: 

```bash
cargo build --target wasm32-wasi --release
wasmedge target/wasm32-wasi/release/noop-nrelay-wasmedge.wasm
```
