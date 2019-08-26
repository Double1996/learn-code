package streamserver

import (
	"io"
	"net/http"
)

func sendErrorMessge(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}
