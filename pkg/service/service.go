package service

import (
	"seo_courses"
	"seo_courses/pkg/repository"
)

type Authorization interface {
	CreateUser(user seo_courses.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(toke string) (int, error)
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

type Service struct {
	Authorization
	Course
	Topic
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Course:        NewCourseService(repos.Course),
		Topic:         NewTopicService(repos.Topic),
	}
}
