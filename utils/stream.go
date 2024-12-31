package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
)

type Message struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

func WriteStreamMessage(w *bufio.Writer, messageType string, input any) error {
	message, _ := json.Marshal(Message{
		Type: messageType,
		Data: input,
	})

	fmt.Fprintf(w, "%s\n", message)

	if err := w.Flush(); err != nil {
		return err
	}

	return nil
}
