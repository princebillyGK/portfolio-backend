package main

import (
	"net/http"

	"princebillygk.portfolio.io/pkg/util"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe("0.0.0.0:"+util.Env("API_PORT", "80"), nil)
}
