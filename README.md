# nostr-noop-relays-wasm
A collection of [no-op Nostr relays](https://github.com/akiomik/nostr-noop-relays?tab=readme-ov-file) runs on WebAssembly.

For instructions for building and running relay impls, see README in respective directories.

## List of Implementations

| Target Wasm Runtime | [WASI Sockets](https://github.com/WebAssembly/wasi-sockets) Impl. | Impl. Details |
|--|--|--|
| [WasmEdge](https://wasmedge.org/docs/) | [wasmedge_wasi_socket](https://github.com/second-state/wasmedge_wasi_socket) | Rust, tokio, tungstenite |
| [Wasmer](https://wasmer.io/) + [WASIX](https://wasix.org/) | WASIX | Rust, tokio, tungstenite |


## What is the no-op relay?
Please refer to [the description](https://github.com/akiomik/nostr-noop-relays?tab=readme-ov-file#what-is-the-noop-relay).
