package recomendation

import (
	"database/sql"
	"math"
	"sort"
)

func collaborativeFiltering(db *sql.DB, studentID int64) ([]int, error) {
	allPreferences, err := GetAllStudentPreferences(db)
	if err != nil {
		return nil, err
	}

	// Получаем предпочтения целевого студента
	targetPreferences, err := GetStudentPreferences(db, studentID)
	if err != nil {
		return nil, err
	}

	// Вычисляем сходство между университетами
	similarUniversities := make(map[int]float64)
	for otherStudentID, preferences := range allPreferences {
		if otherStudentID == int(studentID) {
			continue
		}
		similarity := cosineSimilarity(targetPreferences, preferences)
		for uniID, rating := range preferences {
			similarUniversities[uniID] += rating * similarity
		}
	}

	// Выбираем топ-3 университета
	type universityScore struct {
		ID    int
		Score float64
	}
	var recommendations []universityScore
	for uniID, score := range similarUniversities {
		recommendations = append(recommendations, universityScore{ID: uniID, Score: score})
	}

	// Сортировка и выбор топ-3 рекомендаций
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})
	topRecommendations := make([]int, 0, 3)
	for i := 0; i < 3 && i < len(recommendations); i++ {
		topRecommendations = append(topRecommendations, recommendations[i].ID)
	}

	return topRecommendations, nil
}

// Функция для получения предпочтений студентов
func GetStudentPreferences(db *sql.DB, studentID int64) (map[int]float64, error) {
	preferences := make(map[int]float64)
	rows, err := db.Query("SELECT university_id, rating FROM preferences WHERE student_id = ?", studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var universityID int
		var rating float64
		if err := rows.Scan(&universityID, &rating); err != nil {
			return nil, err
		}
		preferences[universityID] = rating
	}

	return preferences, nil
}

// Функция для получения предпочтений всех студентов
func GetAllStudentPreferences(db *sql.DB) (map[int]map[int]float64, error) {
	allPreferences := make(map[int]map[int]float64)
	rows, err := db.Query("SELECT student_id, university_id, rating FROM preferences")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var studentID, universityID int
		var rating float64
		if err := rows.Scan(&studentID, &universityID, &rating); err != nil {
			return nil, err
		}
		if _, exists := allPreferences[studentID]; !exists {
			allPreferences[studentID] = make(map[int]float64)
		}
		allPreferences[studentID][universityID] = rating
	}

	return allPreferences, nil
}

// Функция для расчета косинусного сходства между двумя университетами
func cosineSimilarity(vec1, vec2 map[int]float64) float64 {
	var dotProduct, magnitudeVec1, magnitudeVec2 float64

	for key, val1 := range vec1 {
		if val2, exists := vec2[key]; exists {
			dotProduct += val1 * val2
		}
		magnitudeVec1 += val1 * val1
	}

	for _, val2 := range vec2 {
		magnitudeVec2 += val2 * val2
	}

	if magnitudeVec1 == 0 || magnitudeVec2 == 0 {
		return 0 // Защита от деления на 0
	}
	return dotProduct / (math.Sqrt(magnitudeVec1) * math.Sqrt(magnitudeVec2))
}

// Функция для получения рекомендаций на основе коллаборативной фильтрации
