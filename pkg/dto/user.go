package dto

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	/*Теги binding используются в Go при работе с валидацией входящих параметров структуры.
	Эти теги используются в том случае, когда вам нужно получить данные из входных параметров
	структуры и проверить их на соответствие определенным правилам, например, что поле не пустое,
	что оно содержит только цифры и т.д. */
}
