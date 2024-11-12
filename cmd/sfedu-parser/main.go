package main

import (
	"github.com/go-pg/pg/v10"
	"log"
	"os"
	"rec-start/pkg/db"
	"rec-start/pkg/parser"
)

func main() {
	// инициализируем конект к бд

	dbc := pg.Connect(&pg.Options{
		Addr:     "localhost:5433",
		User:     "postgres",
		Password: "postgres",
		Database: "rec",
	})
	conn := db.New(dbc)
	v, err := conn.Version()
	die(err)
	log.Println(v)

	p := parser.NewSFEDUParser(dbc)
	p.ParseDirections()
	// инициализируем подключения к менеджеру парсинга

	//вызываем нужный парсинг

}

// die calls log.Fatal if err wasn't nil.
func die(err error) {
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Fatal(err)
	}
}
