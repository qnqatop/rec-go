package recomendation

import (
	"database/sql"
	"html/template"
	"net/http"
	"rec-start/pkg/db"
	"strconv"
	"strings"
)

type Recomendation struct {
	repo *db.RecRepo
}

func NewRecomendation(dbc *sql.DB) *Recomendation {
	return &Recomendation{
		repo: db.NewRepo(dbc),
	}
}

func Rec(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		// Чтение данных из формы
		name := r.FormValue("name")
		age, _ := strconv.Atoi(r.FormValue("age"))
		country := r.FormValue("country")
		city := r.FormValue("city")
		programs := strings.Split(r.FormValue("programs"), ",")

		// Вставка данных абитуриента в базу данных
		result, err := db.Exec("INSERT INTO students (name, age, country, city) VALUES (?, ?, ?, ?)", name, age, country, city)
		if err != nil {
			http.Error(w, "Ошибка сохранения данных", http.StatusInternalServerError)
			return
		}
		studentID, _ := result.LastInsertId()

		// Сохранение предпочтений студента по университетам на основе программ
		for _, program := range programs {
			// Получаем university_id по программе
			var universityID int
			err := db.QueryRow(`
				SELECT university_id 
				FROM university_programs 
				WHERE program_name = ? 
				LIMIT 1`, program).Scan(&universityID)
			if err != nil {
				http.Error(w, "Ошибка при поиске университета для программы: "+program, http.StatusInternalServerError)
				return
			}

			// Вставляем предпочтение в таблицу preferences
			_, err = db.Exec("INSERT INTO preferences (student_id, university_id, rating) VALUES (?, ?, ?)", studentID, universityID, 4.5)
			if err != nil {
				http.Error(w, "Ошибка сохранения предпочтений", http.StatusInternalServerError)
				return
			}
		}

		// Получение рекомендаций на основе коллаборативной фильтрации
		recommendedUniIDs, err := collaborativeFiltering(db, studentID)
		if err != nil {
			http.Error(w, "Ошибка получения рекомендаций", http.StatusInternalServerError)
			return
		}

		// Получаем названия университетов для рекомендаций
		var recommendedUniversities []string
		for _, uniID := range recommendedUniIDs {
			var uniName string
			err := db.QueryRow("SELECT name FROM universities WHERE id = ?", uniID).Scan(&uniName)
			if err == nil {
				recommendedUniversities = append(recommendedUniversities, uniName)
			}
		}

		// Отображаем рекомендации с использованием шаблона
		data := struct {
			Name         string
			Universities []string
		}{
			Name:         name,
			Universities: recommendedUniversities,
		}

		tmpl, err := template.ParseFiles("./pkg/template/recommendations.html")
		if err != nil {
			http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data)
	}
}
