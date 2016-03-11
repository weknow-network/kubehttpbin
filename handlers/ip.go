package handlers

import (
	"net/http"
)

func IP(w http.ResponseWriter, r *http.Request) {
	out := map[string]interface{}{OriginKey: r.RemoteAddr}
	output(out)
}
