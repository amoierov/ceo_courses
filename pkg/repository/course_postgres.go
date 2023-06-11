package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"seo_courses/pkg/dto"
)

type CoursePostgres struct {
	db *sqlx.DB
}

func NewCoursePostgres(db *sqlx.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

func (r *CoursePostgres) Create(course dto.Course) (int, error) {
	/*_, err := r.db.Begin()

	if err != nil {
		return 0, err
	}


	var id int
	createCourseQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", coursesTable)
	row := tx.QueryRow(createCourseQuery, course.Title, course.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersCourseQuery := fmt.Sprintf("INSERT INTO %s (user_id, course_id) VALUES ($1, $2)", userCoursesTable)
	_, err = tx.Exec(createUsersCourseQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()*/
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description, difficulty_level, field_of_activity, duration_days, lang, rating, author) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", coursesTable)
	row := r.db.QueryRow(query, course.Title, course.Description, course.DifficultyLevel, course.FieldOfActivity, course.DurationDays, course.Lang, course.Rating, course.Author) //выполнение запроса и подставление аргументов в плейсхолдеры
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CoursePostgres) Subscribe(userId, courseId int) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, course_id) values ($1, $2)", userCoursesTable)
	_, err := r.db.Exec(query, userId, courseId)
	return err
}
func (r *CoursePostgres) GetAll() ([]dto.Course, error) {
	var courses []dto.Course
	query := fmt.Sprintf("SELECT * FROM %s",
		coursesTable)
	err := r.db.Select(&courses, query)
	return courses, err
}

func (r *CoursePostgres) GetCoursesByIdUser(userId int) ([]dto.Course, error) {
	var list []dto.Course

	query := `SELECT courses.title, courses.description, courses.difficulty_level, courses.field_of_activity, courses.duration_days, courses.lang, courses.rating, courses.author
              FROM courses
              JOIN user_courses ON courses.id = user_courses.course_id
              WHERE user_courses.user_id = $1`
	err := r.db.Select(&list, query, userId)

	return list, err
}

func (r *CoursePostgres) Delete(userId, courseId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.course_id AND ul.user_id=$1 AND ul.course_id=$2",
		coursesTable, userCoursesTable)
	_, err := r.db.Exec(query, userId, courseId)

	return err
}
