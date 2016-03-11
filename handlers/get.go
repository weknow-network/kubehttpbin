package handlers

import (
	"net/http"
)

func Get(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	output(resp, m)
}
