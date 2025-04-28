package main

import(
	"go_blog/database"
	"go_blog/handlers"
	"fmt"
	"log"
	"net/http"
)

func main(){
	database.Connect()
	http.HandleFunc("/api/posts", handlers.GetPosts)
	http.HandleFunc("/api/create", handlers.CreatePost)

	fmt.Println("Server runnning at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
