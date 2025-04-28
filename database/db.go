package database

import(
	"go_blog/models"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)


var db *sql.DB
func Connect(){
	var err error
	dsn := "root:mysql@tcp(127.0.0.1:3307)/blog_db"
	db,err = sql.Open("mysql", dsn)
	if err !=nil{
		log.Fatal(err)
	}
	if err:= db.Ping(); err!=nil{
		log.Fatal(err)

	}
	fmt.Println("Connected to SQL DB")
	}

	func GetAllPosts() ([]models.Post, error){
		rows, err:= db.Query("select id, title, content from posts")
		if err !=nil{
			return nil, err
		}

		defer rows.Close()
		var posts []models.Post
		for rows.Next(){
			var post models.Post 
			if err:= rows.Scan(&post.ID, &post.Title, &post.Content); err!=nil{ 
				return nil, err
		}

			posts = append(posts,post)
		}
		return posts, nil
	}

func CreatePosts( post models.Post) error{
	_,err:= db.Exec("insert into posts(title, content) values(?, ?)", post.Title, post.Content)

	return err
}
