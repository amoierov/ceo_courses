package service

import (
	"seo_courses/pkg/dto"
	"seo_courses/pkg/repository"
)

type CourseService struct {
	repo repository.Course //интерфейс для работы с хранилищем данных.
	// Этот интерфейс описывает методы, которые могут быть использованы для выполнения операций с данными о курсах.
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) Create(course dto.Course) (int, error) {
	return s.repo.Create(course)
}

func (s *CourseService) Subscribe(userId, courseId int) error {
	return s.repo.Subscribe(userId, courseId)
}

func (s *CourseService) GetAll() ([]dto.Course, error) {
	return s.repo.GetAll()
}

func (s *CourseService) GetCoursesByIdUser(userId int) ([]dto.Course, error) {
	return s.repo.GetCoursesByIdUser(userId)
}

func (s *CourseService) Delete(courseId int) error {
	return s.repo.Delete(courseId)
}

func (s *CourseService) UpdateCourse(courseId int, input dto.UpdateCourse) error {
	return s.repo.UpdateCourse(courseId, input)
}

func (s *CourseService) GetAuthors() ([]dto.Author, error) {
	return s.repo.GetAuthors()
}
