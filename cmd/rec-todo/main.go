// main.go
package main

import (
	"database/sql"
	"log"

	"rec-start/pkg/app"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := InitDatabase()
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.Close()

	a := app.New(db)

	if err = a.Run(); err != nil {
		log.Fatal(err)
	}
}

// InitDatabase инициализирует базу данных и создает необходимые таблицы
func InitDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./university_recommendation.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// populateDatabase заполняет базу данных начальными данными
func populateDatabase(db *sql.DB) {
	// Добавление университетов и программ
	universities := []string{"Harvard University", "Stanford University", "MIT"}
	programs := map[string]map[string]float64{
		"Harvard University":  {"Computer Science": 4.9, "Law": 4.7},
		"Stanford University": {"Engineering": 4.8, "Physics": 4.6},
		"MIT":                 {"Engineering": 4.9, "Mathematics": 4.8},
	}

	for _, uni := range universities {
		result, err := db.Exec("INSERT INTO universities (name) VALUES (?)", uni)
		if err != nil {
			log.Println("Университет уже существует:", uni)
			continue
		}
		uniID, _ := result.LastInsertId()
		for program, rating := range programs[uni] {
			_, err := db.Exec("INSERT INTO university_programs (university_id, program_name, rating) VALUES (?, ?, ?)", uniID, program, rating)
			if err != nil {
				log.Println("Ошибка добавления программы:", err)
			}
		}
	}
}
