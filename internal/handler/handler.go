package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func LandingHandler(w http.ResponseWriter, r *http.Request) {
	VerificationDocuments := []string{
		"Passport",
		"Driving Licence",
		"Government ID",
	}
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	tmpl.Execute(w, map[string]interface{}{"VerificationDocuments": VerificationDocuments})
}

func UploadDocumentHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // 32 MB max memory
	if err != nil {
		http.Error(w, "Bad Request: Error parsing form", http.StatusBadRequest)
		return
	}

	documentType := r.FormValue("documentType")
	if documentType == "" {
		fmt.Println("error parsing documentType")
		http.Error(w, "Bad Request: Missing document type", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("documentImage")
	if err != nil {
		fmt.Println("error parsing documentImage")
		http.Error(w, "Bad Request: Error processing document image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Printf("Received: %s\n", documentType)
	// Process the file and documentType as needed...

	w.WriteHeader(http.StatusOK)
}
