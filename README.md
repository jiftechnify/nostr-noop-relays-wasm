# nostr-noop-relays-wasm
A collection of [no-op Nostr relays](https://github.com/akiomik/nostr-noop-relays?tab=readme-ov-file) runs on WebAssembly.

For instructions for building and running relay impls, see README in respective directories.

## List of Implementations

| Target Wasm Runtime | Language | Impl. Details |
|--|--|--|
| [WasmEdge](https://wasmedge.org/docs/) | Rust | tokio, tungstenite(witu both WasmEdge customized, based on [wasmedge_wasi_socket](https://github.com/second-state/wasmedge_wasi_socket)) |
| [Wasmtime]() | Rust  | tokio(with unstable `net` feature support for WASI), tungstenite |
| [Wasmer](https://wasmer.io/) + [WASIX](https://wasix.org/) | Rust | tokio(with WASIX customized), tungstenite |


## What is the no-op relay?
Please refer to [the description](https://github.com/akiomik/nostr-noop-relays?tab=readme-ov-file#what-is-the-noop-relay).
