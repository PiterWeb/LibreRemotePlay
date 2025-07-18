// Websocket server
package websocket

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/coder/websocket"
)

var conns = map[string]*websocket.Conn{}

func SetupWebsocketHandler(serverMux *http.ServeMux) {

	serverMux.HandleFunc("GET /ws", func(w http.ResponseWriter, r *http.Request) {

		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		defer func() {
			err := c.CloseNow()

			if err != nil {
				log.Println(err)
			}
		}()

		// Set the context as needed. Use of r.Context() is not recommended
		// to avoid surprising behavior (see http.Hijacker).
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		wsBroadcast(ctx, r, c)

		c.Close(websocket.StatusNormalClosure, "Client connection closed")

	})

}

func wsBroadcast(ctx context.Context, r *http.Request, ws *websocket.Conn) {
	conns[r.RemoteAddr] = ws

	defer func() {

		// If panic it will recover and liberate resources
		_ = recover()

		for addr := range conns {
			if r.RemoteAddr == addr {
				delete(conns, addr)
				break
			}
		}

	}()

	for {
		typ, reader, err := ws.Reader(ctx)
		if err != nil {
			log.Println(err)
			break
		}

		for addr, con := range conns {

			if r.RemoteAddr == addr {
				continue
			}

			writer, err := con.Writer(ctx, typ)

			if err != nil {
				log.Println(err)
				continue
			}

			_, err = io.Copy(writer, reader)

			if err != nil {
				log.Println(err)
				continue
			}

			// log.Println("Message sended to ", addr)

			err = writer.Close()

			if err != nil {
				log.Println(err)
				continue
			}

		}

	}
}
