package dto

type Question struct {
	ID       int    `json:"id" db:"id"`
	Question string `json:"question" db:"question"`
}

type Answer struct {
	Id              int     `json:"id"`
	UserId          int     `json:"user_id"`
	FormId          int     `json:"form_id"`
	DifficultyLevel int     `json:"difficulty_level" db:"difficulty_level"`
	FieldOfActivity string  `json:"field_of_activity" db:"field_of_activity"`
	DurationDays    int     `json:"duration_days" db:"duration_days"`
	Lang            string  `json:"lang"`
	Rating          float64 `json:"rating"`
	Author          string  `json:"author"`
}

type CreateAnswerRequest struct {
	FormId          int     `json:"form_id"`
	DifficultyLevel int     `json:"difficulty_level" db:"difficulty_level"`
	FieldOfActivity string  `json:"field_of_activity" db:"field_of_activity"`
	DurationDays    int     `json:"duration_days" db:"duration_days"`
	Lang            string  `json:"lang"`
	Rating          float64 `json:"rating"`
	Author          string  `json:"author"`
}

// unused
type RecommendationStruct struct {
	CourseID              int     `json:"course_id" db:"id_course"`
	Title                 string  `json:"title" binding:"required"`
	Description           string  `json:"description"`
	DifficultyLevel       int     `json:"difficulty_level" db:"difficulty_level"`
	AnswerDifficultyLevel int     `json:"difficulty_level" db:"answer_difficulty_level"`
	DurationDays          int     `json:"duration_days" db:"duration_days"`
	AnswerDurationDays    int     `json:"difficulty_level" db:"answer_duration_days"`
	Rating                float64 `json:"rating"`
	AnswerRating          float64 `json:"difficulty_level" db:"answer_rating"`
}

type RecommendationRequest struct {
	CourseID        int    `json:"course_id" db:"course_id"`
	Title           string `json:"title" binding:"required" db:"title"`
	Description     string `json:"description" db:"description"`
	DifficultyLevel string `json:"difficulty_level" db:"difficulty_level"`
	FieldOfActivity string `json:"field_of_activity" db:"field_of_activity"`
	DurationDays    string `json:"duration_days" db:"duration_days"`
	Lang            string `json:"lang"`
	Rating          string `json:"rating",sql:"DEFAULT:0.00"`
	Author          string `json:"author"`
}
