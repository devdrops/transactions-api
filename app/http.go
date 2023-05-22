package app

import (
	"net/http"
)

func StartServer() {
	r := NewRouter()
	http.ListenAndServe(":8080", r)
}
