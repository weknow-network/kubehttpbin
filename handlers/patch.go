package handlers

import (
	"net/http"
)

func Patch(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	output(resp, m)
}
