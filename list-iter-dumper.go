package lf

import (
	"encoding/json"
	"io"
)

func DumpJSON(w io.Writer, it <-chan json.RawMessage) {
	count := 0
	io.WriteString(w, "[")
	for item := range it {
		if count > 0 {
			io.WriteString(w, ",")
		}
		count += 1
		w.Write(item)
	}
	io.WriteString(w, "]")
}

func Dump(w io.Writer, it <-chan json.RawMessage) {
	for item := range it {
		w.Write(item)
		io.WriteString(w, "\n")
	}
}
