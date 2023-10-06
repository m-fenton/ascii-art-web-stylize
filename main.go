package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		showError(w, "404 Page not found.", 404)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./template/index.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Print("HTTP status 500 - Internal Server Errors", err)
			return
		}

		//	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		banner := r.FormValue("banner")
		textbox := r.FormValue("input")

		// errorHandler checks for errors, if no error dedected it'll run ascii-art
		if len(banner) == 0 || len(textbox) == 0 || strings.Contains(textbox, "£") {
			showError(w, "400 Bad Request", 400)
		} else {

			AsciiArt(w, banner, textbox)
			return
		}
	}
}

func showError(w http.ResponseWriter, message string, statusCode int) {
	t, err := template.ParseFiles("./template/errors.html")
	if err == nil {
		w.WriteHeader(statusCode)
		t.Execute(w, message)
		return
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	fmt.Printf("Starting server at http://localhost:8080 ...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("HTTP status 500 - Internal Server Errors")
	}
}
