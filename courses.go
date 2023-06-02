package seo_courses

// Курс
type Course struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
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
