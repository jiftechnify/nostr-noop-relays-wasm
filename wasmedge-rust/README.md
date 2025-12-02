# Nostr no-op relay for WasmEdge (in Rust)

## How to Build and Run
1. Follow the steps [here](https://wasmedge.org/docs/develop/rust/setup/) to setup WasmEdge & Rust toolchain.
2. Build and run the Wasm module by following commands: 

```bash
cargo build --target wasm32-wasip1 --release
wasmedge target/wasm32-wasip1/release/noop-nrelay-wasmedge.wasm
```
