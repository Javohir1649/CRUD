package main

type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"fullName"`
	Email    string `json:"fullName"`
}

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

func main() {
	router := mux.NewRouter()

}
