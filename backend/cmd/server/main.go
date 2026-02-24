package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Will Steinleitner</title>
		</head>
		<body>
			<h1>Website is currently in progress...</h1>
		</body>
	</html>
	`
	w.Write([]byte(html))
}

func main() {
	fmt.Println("Starting server at port 8090...")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}
