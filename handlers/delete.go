package handlers

import (
	"net/http"
)

func Delete(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	output(resp, m)
}
