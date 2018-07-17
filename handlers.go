package main

import (
	"fmt"
	//"encoding/json"
	"glask/db"
	"net/http"
	"path"
	"log"
	"html/template"
)

func Posts(w http.ResponseWriter, r *http.Request){
	posts := db.ShowPosts()
	t, _ := template.ParseFiles("templates/index.gtpl")
 	t.Execute(w, posts)
}

func Post(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		title := r.FormValue("title")
		body := r.FormValue("body")
		log.Println("Creating a new post with title: ", title)
		db.NewPost(title, body)
		//json.NewEncoder(w).Encode(db.NewPost(title))
		http.Redirect(w, r, "/", 301)
	} else {
		log.Println("Invalid HTTP method for a new post")
		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request){
	post := path.Base(r.URL.Path)
	fmt.Println("Deleting post: ", post)
	db.DeletePost(post)
	http.Redirect(w, r, "/", 301)
}

func Home(w http.ResponseWriter, r *http.Request){
	posts := db.ShowPosts()
	t, _ := template.ParseFiles("templates/index.gtpl")
 	t.Execute(w, posts)
}

func About(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("templates/about.gtpl")
 	t.Execute(w, nil)
}