package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"seo_courses/pkg/dto"
)

type FormPostgres struct {
	db *sqlx.DB
}

func NewFormPostgres(db *sqlx.DB) *FormPostgres {
	return &FormPostgres{db: db}
}

func (r *FormPostgres) CreateAnswer(userId int, answer dto.CreateAnswerRequest) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, form_id, difficulty_level, field_of_activity, duration_days, lang, rating, author) VALUES  ($1, $2, $3, $4, $5, $6, $7, $8);", answersTable)
	_, err := r.db.Exec(query, userId, answer.FormId, answer.DifficultyLevel, answer.FieldOfActivity, answer.DurationDays, answer.Lang, answer.Rating, answer.Author)
	return err
}

func (r *FormPostgres) CreateRecommendation(userId int) ([]dto.Course, dto.Answer, error) {
	var list []dto.Course
	var userReferences dto.Answer

	query := fmt.Sprintf(`SELECT courses.id, courses.title, courses.description, courses.field_of_activity, 
       courses.difficulty_level, courses.lang, courses.author, courses.duration_days, courses.rating FROM courses JOIN answers on courses.author = answers.author where answers.id = (
        select max(id) from answers where user_id = $1)`)
	err := r.db.Select(&list, query, userId)
	query = fmt.Sprintf(`SELECT answers.field_of_activity, answers.lang,answers.difficulty_level, answers.duration_days, answers.rating FROM answers where answers.id = (
        select max(id) from answers where user_id = $1)`)
	err = r.db.Get(&userReferences, query, userId)

	return list, userReferences, err
}

func (r *FormPostgres) GetForm(formId int) ([]dto.Question, error) {
	var questions []dto.Question

	query := fmt.Sprintf(`SELECT id, question FROM questions WHERE form_id = $1`)
	rows, err := r.db.Query(query, formId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var q dto.Question
		err = rows.Scan(&q.ID, &q.Question)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}
