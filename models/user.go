package	models

import (
	"config"
	"log"
	"time"
)

type User 	struct {
	Id				int				`json:"id"`
	Login			string		`json:"login"`
	Image			string		`json:"image"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type Users	[]Car

func FindUserById(id int) *Car {
	var user User

	row := config.Db().QueryRow("SELECT * FROM users WHERE id = $1;", id)
	err := row.Scan(&user.Id, &user.Login, &user.Image, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &user
}