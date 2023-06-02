package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"seo_courses"
)

type CoursePostgres struct {
	db *sqlx.DB
}

func NewCoursePostgres(db *sqlx.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

func (r *CoursePostgres) Create(userId int, course seo_courses.Course) (int, error) {
	tx, err := r.db.Begin()
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

	return id, tx.Commit()
}

func (r *CoursePostgres) GetAll(userId int) ([]seo_courses.Course, error) {
	var courses []seo_courses.Course
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.course_id WHERE ul.user_id = $1",
		coursesTable, userCoursesTable)
	err := r.db.Select(&courses, query, userId)
	return courses, err
}

func (r *CoursePostgres) GetById(userId, courseId int) (seo_courses.Course, error) {
	var list seo_courses.Course

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
								INNER JOIN %s ul on tl.id = ul.course_id WHERE ul.user_id = $1 AND ul.course_id = $2`,
		coursesTable, userCoursesTable)
	err := r.db.Get(&list, query, userId, courseId)

	return list, err
}

func (r *CoursePostgres) Delete(userId, courseId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.course_id AND ul.user_id=$1 AND ul.course_id=$2",
		coursesTable, userCoursesTable)
	_, err := r.db.Exec(query, userId, courseId)

	return err
}
