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


func CreatePost(w http.ResponseWriter, r *http.Request){
	var newPosts models.Post
	if err:= json.NewDecoder(r.Body).Decode(&newPosts); err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err:= database.CreatePosts(newPosts)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPosts)
}
