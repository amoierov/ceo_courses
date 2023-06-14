package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"seo_courses/pkg/dto"
)

type TopicPostgres struct {
	db *sqlx.DB
}

func NewTopicPostgres(db *sqlx.DB) *TopicPostgres {
	return &TopicPostgres{db: db}
}

func (r *TopicPostgres) Create(courseId int, topic dto.Topic) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (course_id , title, contents, materials, assignments) values ($1, $2, $3, $4, $5) RETURNING id", topicsTable)

	row := r.db.QueryRow(query, courseId, topic.Title, topic.Content, topic.Materials, topic.Assignments) //выполнение запроса и подставление аргументов в плейсхолдеры
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil

}
