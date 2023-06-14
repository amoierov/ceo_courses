package service

import (
	"seo_courses/pkg/dto"
	"seo_courses/pkg/repository"
)

type Authorization interface {
	CreateUser(user dto.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(toke string) (int, error)
}

type Course interface {
	Create(course dto.Course) (int, error)
	UpdateCourse(courseId int, input dto.UpdateCourse) error
	GetAll() ([]dto.Course, error)
	GetCoursesByIdUser(userId int) ([]dto.Course, error)
	Delete(courseId int) error
	Subscribe(userId, courseId int) error
	GetAuthors() ([]dto.Author, error)
}

type Topic interface {
	Create(courseId int, topic dto.Topic) (int, error)
}

type Form interface {
	UserPreferences(userId int, answer dto.CreateAnswerRequest) error
	CreateRecommendation(userId int) ([]dto.RecommendationRequest, error)
	GetForm(formId int) ([]dto.Question, error)
}

type Service struct {
	Authorization
	Course
	Topic
	Form
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Course:        NewCourseService(repos.Course),
		Topic:         NewTopicService(repos.Topic),
		Form:          NewFormService(repos.Form),
	}
}
