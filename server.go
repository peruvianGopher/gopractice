package main

// ORIGINAL DEMO https://github.com/sausheong/gwp
import (
	"fmt"
	"gopractice/data"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("hi dev!")
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	//mux.HandleFunc("/err", err)

	//mux.HandleFunc("/login", login)
	//mux.HandleFunc("/logout", logout)
	//mux.HandleFunc("/signup", signup)
	//mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	//mux.HandleFunc("/thread/new", newThread)
	//mux.HandleFunc("/thread/create", createThread)
	//mux.HandleFunc("/thread/post", postThread)
	//mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	_ = server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err = session(r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar.html", "templates/index.html")
		} else {
			generateHTML(w, threads, "layout", "private.navbar.html", "templates/index.html")
		}
	}
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(w, "layout", data)
}