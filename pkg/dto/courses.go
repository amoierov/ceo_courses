package dto

// Курс
type Course struct {
	ID              int     `json:"id" db:"id"`
	Title           string  `json:"title" binding:"required" db:"title"`
	Description     string  `json:"description" db:"description"`
	DifficultyLevel int     `json:"difficulty_level" db:"difficulty_level"`
	FieldOfActivity string  `json:"field_of_activity" db:"field_of_activity"`
	DurationDays    int     `json:"duration_days" db:"duration_days"`
	Lang            string  `json:"lang"`
	Rating          float64 `json:"rating",sql:"DEFAULT:0.00"`
	Author          string  `json:"author"`
}

type Author struct {
	CourseId int    `json:"course_id" db:"id"`
	Author   string `json:"author"`
}

type UpdateCourse struct {
	Title           string  `json:"title" binding:"required" db:"title"`
	Description     string  `json:"description" db:"description"`
	DifficultyLevel int     `json:"difficulty_level" db:"difficulty_level"`
	FieldOfActivity string  `json:"field_of_activity" db:"field_of_activity"`
	DurationDays    int     `json:"duration_days" db:"duration_days"`
	Lang            string  `json:"lang"`
	Rating          float64 `json:"rating",sql:"DEFAULT:0.00"`
	Author          string  `json:"author"`
}

// Пользовательские курсы
type UserCourse struct {
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
}

// Топик
type Topic struct {
	ID          int    `json:"id"`
	CourseID    int    `json:"course_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Materials   string `json:"materials"`
	Assignments string `json:"assignments"`
}
