package repository

import (
	"github.com/jmoiron/sqlx"
	"seo_courses/pkg/dto"
)

type Authorization interface {
	CreateUser(user dto.User) (int, error)
	GetUser(username, password string) (dto.User, error)
}

type Course interface {
	Create(course dto.Course) (int, error)
	GetAll() ([]dto.Course, error)
	GetCoursesByIdUser(userId int) ([]dto.Course, error)
	Delete(userId, courseId int) error
	Subscribe(userId, courseId int) error
}

type Topic interface {
	Create(courseId int, topic dto.Topic) (int, error)
}

type Form interface {
	CreateAnswer(userId int, answer dto.CreateAnswerRequest) error
	CreateRecommendation(userId int) ([]dto.Course, dto.Answer, error)
	GetForm(formId int) ([]dto.Question, error)
}
type Repository struct {
	Authorization
	Course
	Topic
	Form
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Course:        NewCoursePostgres(db),
		Topic:         NewTopicPostgres(db),
		Form:          NewFormPostgres(db),
	}
}
