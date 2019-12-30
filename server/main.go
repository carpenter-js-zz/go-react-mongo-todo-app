package main
// https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6

import (
	"fmt"
    "log"
    "net/http"
    "./router"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}