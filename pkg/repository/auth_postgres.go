package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"seo_courses"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user seo_courses.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	/*Переменная row - представляет собой результат выполнения запроса query с передачей ему аргументов
	user.Name, user.Username, user.Password.
	Эта переменная является указателем на объект типа Row, который может содержать одну строку результата.
	Далее, используется метод Scan() этой переменной для извлечения значения из строки результата и записи его в переменную id.
	Ошибки, возникающие при выполнении метода Scan(), обрабатываются в блоке if-else.*/
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password) //выполнение запроса и подставление аргументов в плейсхолдеры
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (seo_courses.User, error) {
	var user seo_courses.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
