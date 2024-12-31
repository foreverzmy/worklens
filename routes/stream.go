package routes

import (
	"bufio"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

func ResponseStreamHandler(ctx *fasthttp.RequestCtx) {
	// Send the response in chunks and wait for a second between each chunk.
	ctx.SetBodyStreamWriter(func(w *bufio.Writer) {
		for i := 0; i < 10; i++ {
			fmt.Fprintf(w, "this is a message number %d", i)

			// Do not forget flushing streamed data to the client.
			if err := w.Flush(); err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	})
}
