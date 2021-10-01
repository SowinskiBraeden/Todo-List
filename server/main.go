package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SowinskiBraeden/Todo-List/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting the Server on port 9000...")

	log.Fatal(http.ListenAndServe(":9000", r))
}
