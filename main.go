package main

import "fmt"

func main() {
	server := NewServer(":3000")
	var x int32
	fmt.Printf("%d", &x)
	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/api", HandleApi)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("GET", "/home", server.AddMiddleware(HandleHome, CheckAuth(), Logger()))
	server.Listen()
}
