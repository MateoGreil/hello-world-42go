package	models

import (
	"github.com/mateogreil/hello-world-42go/config"
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

type Users	[]User

func FindUserById(id int) *User {
	var user User

	row := config.Db().QueryRow("SELECT users.id, users.login, users.image, users.created_at, users.updated_at FROM users WHERE id = $1;", id)
	err := row.Scan(&user.Id, &user.Login, &user.Image, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &user
}