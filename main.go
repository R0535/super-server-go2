package main

func main() {
	server := NewServer(":3000")

	server.Handle("GET", "/", HandleRoot)
	server.Handle("GET", "/api", HandleApi)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	//to send to the git the new version
	server.Handle("GET", "/home", server.AddMiddleware(HandleHome, CheckAuth(), Logger()))
	server.Listen()
}
