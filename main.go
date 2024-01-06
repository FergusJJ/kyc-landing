package main

import (
	"fmt"
	"net/http"

	"github.com/FergusJJ/kyc-landing/internal/handler"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler.LandingHandler)
	http.HandleFunc("/upload-document", handler.UploadDocumentHandler)
	err := http.ListenAndServe(":8000", nil)
	defer fmt.Println(err)
}
