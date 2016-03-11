package handlers

import (
	"net/http"
)

func Headers(w http.ResponseWriter, r *http.Request) {
	out := make(map[string]interface{})
	hdrs := r.Header
	for key, vals := range hdrs {
		out[key] = vals[0]
	}
	output(w, out)
}
