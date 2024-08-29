package markdownnoteapp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gorilla/mux"
)

/**
Scope of app
------------
-> upload markdown file [✅]
-> check the grammer
-> shows all notes      [✅]
-> render in html       [✅]

Aim
---
The goal of this project is to help you learn how to handle file uploads in a RESTful API,
parse and render markdown files using libraries, and check the grammar of the notes

Features
--------
You’ll provide an endpoint to check the grammar of the note.
You’ll also provide an endpoint to save the note that can be passed in as Markdown text.
Provide an endpoint to list the saved notes (i.e. uploaded markdown files).
Return the HTML version of the Markdown note (rendered note) through another endpoint.


Requirement gethering & things to explore
-----------------------------------------
-> markdown internals
-> working with markdown in golang
-> file handling in golang

API Endpoints
-------------
/checkgrammer POST
/upload 	  POST
/notes 		  GET
/notes/{id}   GET

**/

// Upload Markdown file
func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, _ := r.FormFile("file")
	defer file.Close()

	// check for file is md file or not or else return invalid
	fileExtension := strings.ToLower(filepath.Ext(header.Filename))

	if fileExtension != ".md" {
		log.Println("Invalid File Type")
		json.NewEncoder(w).Encode("Invalid File Type, upload only markdown file")
		return
	}

	dst, _ := os.Create(filepath.Join("./advanced/markdown-note-app/notes", header.Filename))
	defer dst.Close()

	io.Copy(dst, file)

	json.NewEncoder(w).Encode("File uploaded success")
}

// List all notes: return the list of the filename of all uploaded notes
func listAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	notes, err := os.ReadDir("./advanced/markdown-note-app/notes")
	if err != nil {
		log.Println("Error while reading directory: ", err)
		json.NewEncoder(w).Encode("Error while fetching notes")
		return
	}

	var notesListResponse []string
	for _, note := range notes {
		name := note.Name()
		name = name[:len(name)-3]
		notesListResponse = append(notesListResponse, name)
	}

	json.NewEncoder(w).Encode(notesListResponse)
}

// Check Grammer: Parse the file and return the validity with error at places
func checkGrammerHandler(w http.ResponseWriter, r *http.Request) {

}

// View in HTML
func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	filePath := "./advanced/markdown-note-app/notes/" + args["name"] + ".md"
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error while reading file: ", err)
		json.NewEncoder(w).Encode("Error while fetching note")
		return
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(file)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	w.Header().Set("Content-Type", "text/html")
	w.Write(markdown.Render(doc, renderer))
}

func MarkdownAppStart() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", fileUploadHandler)
	r.HandleFunc("/notes", listAllNotesHandler)
	r.HandleFunc("/notes/{name}", getNoteHandler)
	r.HandleFunc("/checkGrammer", checkGrammerHandler)
	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Println("Err: ", err)
	}
}
