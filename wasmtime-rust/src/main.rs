use std::{net::SocketAddr, os::fd::FromRawFd};

use futures::{SinkExt, StreamExt};
use serde_json::json;
use tokio::net::{TcpListener, TcpStream};
use tokio_tungstenite::{accept_async as ws_accept_async, tungstenite::Message, WebSocketStream};
use valq::query_value;

struct Conn(WebSocketStream<TcpStream>, SocketAddr);

#[tokio::main(flavor = "current_thread")]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let listener = unsafe { std::net::TcpListener::from_raw_fd(3) };
    let listener = TcpListener::from_std(listener)?;

    loop {
        let (tcp_strm, raddr) = listener.accept().await?;
        let ws = ws_accept_async(tcp_strm).await?;

        tokio::spawn(serve_conn(Conn(ws, raddr)));
    }
}

async fn serve_conn(conn: Conn) {
    let Conn(mut ws, raddr) = conn;
    println!("connected from {raddr}");

    while let Some(msg) = ws.next().await {
        match msg {
            Ok(m) if m.is_text() => {
                println!("[{raddr}] received message: {m}");
                if let Some(resp) = handle_nostr_msg(m.to_text().unwrap()) {
                    if let Err(e) = ws.send(Message::Text(resp)).await {
                        println!("[{raddr} failed to respond: {e}]");
                    }
                }
            }
            Ok(_) => println!("[{raddr}] received non-text message"),
            Err(e) => println!("[{raddr}] error: {e}"),
        }
    }
    println!("[{raddr}] websocket closed");
}

fn handle_nostr_msg(raw: &str) -> Option<String> {
    let msg: serde_json::Value = serde_json::from_str(raw).ok()?;

    let resp = match query_value!(msg[0] -> str)? {
        "EVENT" => Some(json!(["OK", query_value!(msg[1].id -> str)?, true, ""])),
        "REQ" => Some(json!(["EOSE", query_value!(msg[1] -> str)?])),
        _ => None,
    };
    resp.map(|v| v.to_string())
}
