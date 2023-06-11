package service

import (
	"math"
	"seo_courses/pkg/dto"
	"seo_courses/pkg/repository"
	"strconv"
)

type FormService struct {
	formRepo repository.Form
}

func (s *FormService) GetForm(formId int) ([]dto.Question, error) {
	return s.formRepo.GetForm(formId)
}

func NewFormService(repo repository.Form) *FormService {
	return &FormService{formRepo: repo}
}

func (s *FormService) UserPreferences(userId int, answer dto.CreateAnswerRequest) error {
	return s.formRepo.CreateAnswer(userId, answer)
}

func (s *FormService) CreateRecommendation(userId int) ([]dto.RecommendationRequest, error) {
	listCourses, userReferences, err := s.formRepo.CreateRecommendation(userId)
	if err != nil {
		return nil, err
	}
	recommendedCourses := recommendCourses(listCourses, userReferences)
	res := GenerateRecommendations(recommendedCourses, userReferences)
	return res, nil
}

func cosineSimilarity(rec dto.Course, user dto.Answer) float64 {
	// вектор курса
	courseVec := []float64{float64(rec.DifficultyLevel), float64(rec.DurationDays), rec.Rating}
	// вектор предпочтений пользователя
	userVec := []float64{float64(user.DifficultyLevel), float64(user.DurationDays), user.Rating}

	// dot product
	dotProduct := 0.0
	// длина вектора курса
	courseNorm := 0.0
	// длина вектора предпочтений пользователя
	userNorm := 0.0
	// итерация по элементам векторов
	for i := 0; i < len(courseVec); i++ {
		dotProduct += courseVec[i] * userVec[i]
		courseNorm += math.Pow(courseVec[i], 2)
		userNorm += math.Pow(userVec[i], 2)
	}

	// вычисление длин векторов
	courseNorm = math.Sqrt(courseNorm)
	userNorm = math.Sqrt(userNorm)

	// вычисление косинусного расстояния
	cosSim := dotProduct / (courseNorm * userNorm)

	return cosSim
}

func levenshtein(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	// Создаем двумерный массив для хранения результата сравнения
	// между двумя строками
	matrix := make([][]int, m+1)
	for i := range matrix {
		matrix[i] = make([]int, n+1)
	}

	// Заполняем первую строку и столбец матрицы
	for i := 0; i <= m; i++ {
		matrix[i][0] = i
	}

	for j := 0; j <= n; j++ {
		matrix[0][j] = j
	}

	// Заполняем оставшуюся часть матрицы по правилам динамического программирования
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			var cost int
			if s1[i-1] == s2[j-1] {
				cost = 0
			} else {
				cost = 1
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,
				matrix[i][j-1]+1,
				matrix[i-1][j-1]+cost,
			)
		}
	}

	// Возвращаем значение наименьшего количества операций из правого нижнего угла матрицы
	return matrix[m][n]
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}

func recommendCourses(recs []dto.Course, user dto.Answer) []dto.Course {
	var recommended []dto.Course

	// перебор всех курсов и расчет косинусного расстояния
	for _, rec := range recs {
		// проверить, если у курса совпадают требования пользователя
		if rec.DifficultyLevel == user.DifficultyLevel &&
			rec.DurationDays == user.DurationDays &&
			rec.Rating == user.Rating {
			recommended = append(recommended, rec)
			continue
		}
		// Вычисляем расстояние Левенштейна между значениями поля FieldOfActivity
		levField := levenshtein(rec.FieldOfActivity, user.FieldOfActivity)
		if levField > 3 { // Если расстояние больше порога, то идем дальше
			continue
			// Вычисляем расстояние Левенштейна между значениями поля Lang
		}
		levLang := levenshtein(rec.Lang, user.Lang)
		if levLang > 2 { // Если расстояние больше порога, то идем дальше
			continue
		}

		// расчет косинусного расстояния с предпочтениями пользователя
		cosSim := cosineSimilarity(rec, user)
		if cosSim >= 0.9 {
			recommended = append(recommended, rec)
		}
	}
	return recommended
}

func GenerateRecommendations(recommendedCourses []dto.Course, answer dto.Answer) []dto.RecommendationRequest {
	var recommendations []dto.RecommendationRequest

	for _, course := range recommendedCourses {
		recommendation := dto.RecommendationRequest{
			CourseID:        course.ID,
			Title:           course.Title,
			Description:     course.Description,
			FieldOfActivity: course.FieldOfActivity,
			Lang:            course.Lang,
			Author:          course.Author,
		}

		if course.DifficultyLevel > answer.DifficultyLevel {
			recommendation.DifficultyLevel = "hard - " + strconv.Itoa(course.DifficultyLevel)
		} else if course.DifficultyLevel == answer.DifficultyLevel {
			recommendation.DifficultyLevel = "optimal - " + strconv.Itoa(course.DifficultyLevel)
		} else {
			recommendation.DifficultyLevel = "easy - " + strconv.Itoa(course.DifficultyLevel)
		}

		if course.DurationDays > answer.DurationDays {
			recommendation.DurationDays = "too long - " + strconv.Itoa(course.DurationDays)
		} else if course.DurationDays == answer.DurationDays {
			recommendation.DurationDays = "optimal - " + strconv.Itoa(course.DurationDays)
		} else {
			recommendation.DurationDays = "shorter - " + strconv.Itoa(course.DurationDays)
		}

		if course.Rating > 5 && course.Rating < 8 {
			recommendation.Rating = "medium rating(yellow) - " + strconv.FormatFloat(course.Rating, 'f', 2, 64)
		} else if course.Rating >= 8 {
			recommendation.Rating = "high rating(green) - " + strconv.FormatFloat(course.Rating, 'f', 2, 64)
		} else {
			recommendation.Rating = "low rating(red) - " + strconv.FormatFloat(course.Rating, 'f', 2, 64)
		}

		recommendations = append(recommendations, recommendation)
	}

	return recommendations
}

func findClosestStruct(structs []dto.RecommendationStruct) dto.RecommendationStruct {
	closestDiff := math.MaxInt32
	closestDur := math.MaxInt32
	closestRating := math.MaxFloat64
	var closestStruct dto.RecommendationStruct
	// перебираем все структуры, переданные в функцию
	for _, s := range structs {
		// вычисляем значения отклонений (разниц) между данными рекомендации и этих же данных, но со знаком "answer"
		diff := int(math.Abs(float64(s.DifficultyLevel - s.AnswerDifficultyLevel)))
		dur := int(math.Abs(float64(s.DurationDays - s.AnswerDurationDays)))
		rating := math.Abs(s.Rating - s.AnswerRating)
		// если данная структура более близка к ответу, чем ближайший найденный answer,
		//то запоминаем данную структуру, как новый ближайший answer
		if diff < closestDiff || (diff == closestDiff && (dur < closestDur || (dur == closestDur && rating < closestRating))) {
			closestDiff = diff
			closestDur = dur
			closestRating = rating
			closestStruct = s
		}
	}

	return closestStruct
}
