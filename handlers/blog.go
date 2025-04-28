package handlers

import(
	"go_blog/database"
	"go_blog/models"
	"encoding/json"
	"net/http"
)


func GetPosts(w http.ResponseWriter, r *http.Request){
	posts, err:= database.GetAllPosts()
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}


func CreatePosts(w http.ResponseWriter, r *http.Request){
	var newPosts= models.Post
	if err:= json.NewDecoder(r.body).Decode(&newPosts); err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
