# Nostr no-op relay for Wasmtime (in Rust, on top of WASI 0.2)

## How to Build
Prerequisite: Rust programming environment (`rustc`, `cargo`, `rustup`) has been installed.

1. Add `wasm32-wasip2` target to the Rust compiler.

```bash
rustup target add wasm32-wasip2
```

2. Build the Wasm module. 

```bash
cargo build --target wasm32-wasip2 --release
```

## How to Run

You can run the built Wasm module on any Wasm runtimes supporting WASI 0.2. As of today, `wasmtime` and `jco` are the only runtimes that support it, among major Wasm runtimes.

### Run on wasmtime

Install wasmtime by following [the instruction](https://wasmtime.dev/).

```bash
wasmtime run -S inherit-network=y  target/wasm32-wasip2/release/noop-nrelay-wstd.wasm
```

### Run on jco
 
Install jco by following [the instruction](https://github.com/bytecodealliance/jco?tab=readme-ov-file#quickstart).

```bash
jco run target/wasm32-wasip2/release/noop-nrelay-wstd.wasm
```
