package service

import (
	"seo_courses"
	"seo_courses/pkg/repository"
)

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) Create(userId int, course seo_courses.Course) (int, error) {
	return s.repo.Create(userId, course)
}

func (s *CourseService) GetAll(userId int) ([]seo_courses.Course, error) {
	return s.repo.GetAll(userId)
}

func (s *CourseService) GetById(userId, courseId int) (seo_courses.Course, error) {
	return s.repo.GetById(userId, courseId)
}

func (s *CourseService) Delete(userId, courseId int) error {
	return s.repo.Delete(userId, courseId)
}
