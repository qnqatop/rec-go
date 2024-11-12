package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Функция для предзаполнения базы данных
func populateDatabase(db *sql.DB) {
	// Добавление университетов
	universities := []string{"Harvard University", "Stanford University", "MIT", "Oxford University", "Cambridge University"}
	for _, uni := range universities {
		_, err := db.Exec("INSERT INTO universities (name) VALUES (?)", uni)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Добавление программ для университетов
	universityPrograms := []struct {
		UniversityName string
		Programs       map[string]float64
	}{
		{"Harvard University", map[string]float64{"Computer Science": 4.9, "Law": 4.7, "Medicine": 4.8}},
		{"Stanford University", map[string]float64{"Computer Science": 4.8, "Engineering": 4.7, "Physics": 4.6}},
		{"MIT", map[string]float64{"Engineering": 4.9, "Computer Science": 4.9, "Mathematics": 4.8}},
		{"Oxford University", map[string]float64{"Law": 4.8, "Literature": 4.7, "Politics": 4.6}},
		{"Cambridge University", map[string]float64{"Engineering": 4.8, "Biology": 4.7, "Physics": 4.6}},
	}

	for _, uniProg := range universityPrograms {
		var uniID int
		err := db.QueryRow("SELECT id FROM universities WHERE name = ?", uniProg.UniversityName).Scan(&uniID)
		if err != nil {
			log.Fatal(err)
		}

		for prog, rating := range uniProg.Programs {
			_, err := db.Exec("INSERT INTO university_programs (university_id, program_name, rating) VALUES (?, ?, ?)", uniID, prog, rating)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// Добавление студентов
	students := []struct {
		Name    string
		Age     int
		Country string
		City    string
	}{
		{"Alice", 18, "USA", "New York"},
		{"Bob", 19, "USA", "San Francisco"},
		{"Charlie", 18, "UK", "London"},
	}

	for _, student := range students {
		_, err := db.Exec("INSERT INTO students (name, age, country, city) VALUES (?, ?, ?, ?)", student.Name, student.Age, student.Country, student.City)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func pre() {
	// Создание базы данных и предзаполнение
	db, err := sql.Open("sqlite3", "./university_recommendation.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Заполнение базы данных
	populateDatabase(db)

	fmt.Println("Database has been populated with universities, programs, and students.")
}
