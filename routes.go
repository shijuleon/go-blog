package main

import (
	"net/http"
)

func InitRoutes(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/posts/new", Post)
	http.HandleFunc("/posts", Posts)
	http.HandleFunc("/about", About)
	http.HandleFunc("/delete/", Delete)
}