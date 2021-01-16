package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/alacine/deptool/defs"
)

func sendNormalResponse(w http.ResponseWriter, sc int, resp defs.Resp) {
	w.WriteHeader(sc)
	resp.State = state
	respStr, _ := json.Marshal(resp)
	io.WriteString(w, string(respStr))
}

func sendErrorResponse(w http.ResponseWriter, sc int, resp defs.Resp, err error) {
	w.WriteHeader(sc)
	resp.State = state
	resp.Message += err.Error()
	respStr, _ := json.Marshal(resp)
	io.WriteString(w, string(respStr))
}
