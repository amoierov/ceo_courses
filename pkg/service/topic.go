package service

import (
	"seo_courses"
	"seo_courses/pkg/repository"
)

type TopicService struct {
	repo       repository.Topic
	courseRepo repository.Course
}

func NewTopicService(repo repository.Topic) *TopicService {
	return &TopicService{repo: repo}
}

func (s *TopicService) Create(courseId int, topic seo_courses.Topic) (int, error) {
	return s.repo.Create(courseId, topic)
}
