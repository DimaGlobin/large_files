package main

import (
	"net/http"

	"github.com/DimaGlobin/large_files/controller"
)

func setup_routes() {
	http.HandleFunc("/upload", controller.ChunkUploadHandler)
}

func main() {

	setup_routes()

	http.ListenAndServe(":8080", nil)
}
