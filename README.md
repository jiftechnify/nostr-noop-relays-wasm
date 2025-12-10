# nostr-noop-relays-wasm
A collection of [no-op Nostr relays](https://github.com/akiomik/nostr-noop-relays?tab=readme-ov-file) run on WebAssembly.

For instructions for building and running relay impls, see README in respective directories.

## List of Implementations

| Target Wasm Runtime | WASI ver. | Language | Impl. Details |
|--|--|--|--|
| [WasmEdge](https://wasmedge.org/docs/) | 0.1 | Rust | tokio, tungstenite(with both WasmEdge customized, based on [wasmedge_wasi_socket](https://github.com/second-state/wasmedge_wasi_socket)) |
| WasmEdge | 0.1 | Go | [stealthrocket/net](https://github.com/stealthrocket/net) for WASI Sockets support, [nhooyr.io/websocket](https://github.com/nhooyr/websocket)  |
| [Wasmer](https://wasmer.io/) + [WASIX](https://wasix.org/) | 0.1 | Rust | tokio(with WASIX customized), tungstenite |
| [Wasmtime](https://wasmtime.dev/) | 0.1 | Rust | tokio(with unstable `net` feature support for WASI), tungstenite |
| Wasmtime/[Jco](https://github.com/bytecodealliance/jco)[^1] | 0.2 | Rust | [wstd](https://github.com/bytecodealliance/wstd), tungstenite([patched](https://github.com/jiftechnify/wstd-tungstenite) for wstd) |

[^1]: Theoretically, it can be run on all Wasm runtimes supporting WASI 0.2. As of today, Wasmtime and Jco are the only runtimes that support it, among major Wasm runtimes, though.


## What is the no-op relay?
Please refer to [the description](https://github.com/akiomik/nostr-noop-relays?tab=readme-ov-file#what-is-the-noop-relay).
