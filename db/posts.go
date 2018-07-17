package db

import (
	"database/sql"
	//_ "github.com/mattn/go-sqlite3"
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	 "time"
)

// parse the json file sp

type Post struct {
	Pid int `json:"pid"`
	Title string `json:"title"`
	Body string `json:"body"`
	Time time.Time `json:"time"`
}

// func getPosts(filename string) *Post {
// 	var post Post
// 	content, e := ioutil.ReadFile(filename)
// 	if e != nil {
// 		fmt.Printf("File error") 
// 	}
// 	post.Body = string(content)
// 	return &post
// }

const dbname = "blog.db" 
var db, err = sql.Open("sqlite3", dbname)

func insertPosts(title string, body string) Post {
	stmt, err := db.Prepare("INSERT INTO posts(title, body, time) values (?,?,?)")
	checkErr(err)

	var post Post
	post.Title = title
	post.Body = body
	post.Time = time.Now()
	
	_, err = stmt.Exec(post.Title, post.Body, post.Time)
	checkErr(err)
	return post
}

func queryPosts() []Post {
	var	posts []Post

	rows, err := db.Query("select pid, title, body, time from posts")
	checkErr(err)

	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Body, &post.Time)
		checkErr(err)
		posts = append(posts, post)
	}
	
	return posts
}

func DeletePost(pid string){
	stmt, err := db.Prepare("delete from posts where pid = ?")
	checkErr(err)

	_, err = stmt.Exec(pid)
	checkErr(err)
}

func NewPost(title string, body string) Post{
	return insertPosts(title, body)
}

func ShowPosts() []Post{
	return queryPosts()
}