package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alacine/deploy/defs"
)

func sendNormalResponse(w http.ResponseWriter, sc int, resp defs.Resp) {
	w.WriteHeader(sc)
	respStr, _ := json.Marshal(resp)
	io.WriteString(w, string(respStr))
}

func sendErrorResponse(w http.ResponseWriter, sc int, resp defs.Resp, err error) {
	w.WriteHeader(sc)
	resp.Message += err.Error()
	respStr, _ := json.Marshal(resp)
	io.WriteString(w, string(respStr))
}
