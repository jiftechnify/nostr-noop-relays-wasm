# Nostr no-op relay for Wasmer + WASIX

## How to Build and Run (Single-threaded version)
Prerequisite: Rust programming environment (`rustc`, `cargo`, `rustup`) has been installed.

1. Follow the instrcution [here](https://docs.wasmer.io/install) to install Wasmer CLI.
2. Follow the instruction [here](https://wasix.org/docs/language-guide/rust/installation) to install `cargo-wasix`.
3. Build the Wasm module by: 

```bash
cargo wasix build --release
```

4. Run the Wasm module by:

```bash
wasmer run --net target/wasm32-wasmer-wasi/release/noop-nrelay-wasix.wasm
```

## How to Build and Run (Multi-threaded version)
1. Complete steps 1.-3. for the single-threaded version.
2. Run the Wasm module by:

```bash
wasmer run --net --enable-threads target/wasm32-wasmer-wasi/release/noop-nrelay-wasix_multi-thread.wasm
```
