package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"noop-nrelay-wasmedge-go/nostr"

	"github.com/stealthrocket/net/wasip1"
	"nhooyr.io/websocket"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}
		defer c.CloseNow()

		fmt.Println("connected from ", r.RemoteAddr)
		for {
			msgType, recv, err := c.Read(context.TODO())
			if err != nil {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					fmt.Printf("[%s] websocket closed\n", r.RemoteAddr)
					return
				}
				fmt.Printf("[%s] error: %v\n", r.RemoteAddr, err)
				return
			}
			if msgType != websocket.MessageText {
				fmt.Printf("[%s] received non-text message\n", r.RemoteAddr)
				continue
			}

			fmt.Printf("[%s] received message: %s", r.RemoteAddr, string(recv))
			c2r := nostr.ParseC2RMsg(recv)
			if c2r == nil {
				continue
			}
			r2c := handleC2RMsg(c2r)
			if r2c == nil {
				continue
			}
			if err := writeR2CMsg(c, r2c); err != nil {
				fmt.Printf("[%s] failed to respond: %v\n", r.RemoteAddr, err)
				return
			}
		}
	})

	listener, err := wasip1.Listen("tcp", "127.0.0.1:9876")
	if err != nil {
		return
	}
	if http.Serve(listener, nil); err != nil {
		return
	}
}

func handleC2RMsg(c2r nostr.C2R) nostr.R2C {
	switch c2r := c2r.(type) {
	case *nostr.C2R_REQ:
		return &nostr.R2C_EOSE{SubID: c2r.SubID}
	case *nostr.C2R_EVENT:
		return &nostr.R2C_OK{EventID: c2r.Event.ID, OK: true, Reason: ""}
	}
	return nil
}

func writeR2CMsg(ws *websocket.Conn, r2c nostr.R2C) error {
	w, err := ws.Writer(context.TODO(), websocket.MessageText)
	if err != nil {
		return err
	}
	defer w.Close()
	if err := json.NewEncoder(w).Encode(r2c); err != nil {
		return err
	}
	return nil
}
