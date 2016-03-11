package handlers

import (
	"net/http"
)

func Put(resp http.ResponseWriter, req *http.Request) {
	m := getBase(req)
	m[DataKey] = string(readBody(req))
	output(resp, m)
}
