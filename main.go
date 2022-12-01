package main

import ( 
	"fmt"
	"html/template"
	"net/http"
)

const port = 8080

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/game", Game)
	
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
