package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const templatePath = "./template"
const layoutPath = templatePath + "/layout.html"

var (
	indexTemplate = template.Must(template.ParseFiles(layoutPath, templatePath+"/index.html"))
)

func main() {
	http.HandleFunc("/", IndexHandler)
	fmt.Println("Server is listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"PageTitle": "Home Page",
		"Text":      "Hello w",
	})
}
