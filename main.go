package main

import (
	"./users"
	"log"
	"net/http"
)

func main() {
	repo := &users.Repository{}
	_, err := repo.InitDB()
	if err != nil {
		log.Panic(err)
	}

	users.NewRouter()
	http.ListenAndServe(":8080", nil)
}
