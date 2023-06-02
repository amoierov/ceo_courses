package repository

import (
	"github.com/jmoiron/sqlx"
	"seo_courses"
)

type Authorization interface {
	CreateUser(user seo_courses.User) (int, error)
	GetUser(username, password string) (seo_courses.User, error)
}

type Course interface {
	Create(userId int, course seo_courses.Course) (int, error)
	GetAll(userId int) ([]seo_courses.Course, error)
	GetById(userId, courseId int) (seo_courses.Course, error)
	Delete(userId, courseId int) error
}

type Topic interface {
	Create(courseId int, topic seo_courses.Topic) (int, error)
}

type Repository struct {
	Authorization
	Course
	Topic
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Course:        NewCoursePostgres(db),
		Topic:         NewTopicPostgres(db),
	}
}
