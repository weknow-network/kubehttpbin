package handlers

import (
	"net/http"
)

func Head(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	output(resp, m)
}
