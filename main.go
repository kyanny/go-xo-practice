package main

import (
	"fmt"
	"log"
	"database/sql"
	"github.com/kyanny/go-xo-practice/models"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=dvdrental sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	actor, err := models.ActorByActorID(models.XODB(db), 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", actor)
}
