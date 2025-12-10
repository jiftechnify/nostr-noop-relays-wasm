use futures::{SinkExt, StreamExt};
use serde_json::json;
use valq::query_value;
use wstd::iter::AsyncIterator;
use wstd::net::{TcpListener, TcpStream};
use wstd_tungstenite::{WebSocketStream, accept_async as ws_accept_async, tungstenite::Message};

struct Conn(WebSocketStream<TcpStream>, String);

#[wstd::main]
async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    let listener = TcpListener::bind("127.0.0.1:9876").await?;

    let mut incoming_conns = listener.incoming();
    while let Some(stream) = incoming_conns.next().await {
        let stream = stream?;
        let raddr = stream.peer_addr()?.clone();
        let ws = ws_accept_async(stream).await?;
        wstd::runtime::spawn(serve_conn(Conn(ws, raddr))).detach();
    }
    Ok(())
}

async fn serve_conn(conn: Conn) {
    let Conn(mut ws, raddr) = conn;
    println!("connected from {raddr}");

    while let Some(msg) = ws.next().await {
        match msg {
            Ok(m) if m.is_text() => {
                println!("[{raddr}] received message: {m}");
                #[allow(clippy::collapsible_if)]
                if let Some(resp) = handle_nostr_msg(m.to_text().unwrap()) {
                    if let Err(e) = ws.send(Message::Text(resp.into())).await {
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
