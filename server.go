package main

import (
	// "crypto/md5"
	// "fmt"
	// "html/template"
	// "io"
	"net/http"
	"log"
	// "os"
	// "strconv"
	// "time"
)

// func sayHello(w http.ResponseWriter, r *http.Request){
// 	r.ParseForm()
// 	fmt.Println(r.Form)
// 	for k, v := range r.Form {
// 		fmt.Println("key: ", k)
// 		fmt.Println("value: ", v)
// 	}
// 	fmt.Fprintf(w, "Hello Shiju")
// }

// func home(w http.ResponseWriter, r *http.Request){
// 	token := 1212
// 	t, _ := template.ParseFiles("templates/index.gtpl")
// 	t.Execute(w, token)
// }

// func login(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("method: ", r.Method)
// 	if r.Method == "GET" {
// 		crutime := time.Now().Unix()
// 		h := md5.New()
// 		io.WriteString(h, strconv.FormatInt(crutime, 10))
// 		token := fmt.Sprintf("%x", h.Sum(nil))
// 		t, _ := template.ParseFiles("templates/login.gtpl")
// 		t.Execute(w, token)
// 	} else {
// 		r.ParseForm()
// 		fmt.Println("username: ", r.Form["username"])
// 	}
// }

// func upload(w http.ResponseWriter, r *http.Request){
// 	fmt.Println("method: ", r.Method)
// 	if r.Method == "GET" {
// 		crutime := time.Now().Unix()
// 		h := md5.New()
// 		io.WriteString(h, strconv.FormatInt(crutime, 10))
// 		token := fmt.Sprint("%x", h.Sum(nil))
// 		t, _ := template.ParseFiles("templates/upload.gtpl")
// 		t.Execute(w, token)
// 	} else {
// 		r.ParseMultipartForm(32 << 20)
// 		file, handler, err := r.FormFile("uploadfile")
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		defer file.Close()
// 		fmt.Fprintf(w, "%v", handler.Header)
// 		f, err := os.OpenFile("./test/"+handler.Filename[10:], os.O_WRONLY|os.O_CREATE, 0666)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		defer f.Close()
// 		io.Copy(f, file)
//  	}
// }

func main(){
	InitRoutes()
	log.Println("Starting HTTP server at localhost:7000")
	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}