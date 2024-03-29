package	models

import (
	"github.com/mateogreil/hello-world-42go/config"
	"github.com/mateogreil/hello-world-42go/utils"
	"log"
	"time"
)

type User 	struct {
	Id				int								`json:"id"`
	Login			string						`json:"login"`
	Email			string						`json:"email"`
	CreatedAt	time.Time					`json:"created_at"`
	UpdatedAt	time.Time					`json:"updated_at"`
	Image			utils.NullString	`json:"image"`
	PoolYear	utils.NullString	`json:"pool_year"`
	PoolMonth	utils.NullString	`json:"pool_month"`
	Kind			string						`json:"kind"`
	Status		utils.NullString	`json:"status"`
}

type Users	[]User

func AllUser(page int) *Users {
	var limitPerPage = 10
	var users Users

	rows, err := config.Db().Query(
		"SELECT users.id, " +
			"users.login, " +
			"users.email, " +
			"users.created_at, " +
			"users.updated_at, " +
			"users.image, " +
			"users.pool_year, " +
			"users.pool_month, " +
			"users.kind, " +
			"users.status " +
		"FROM users " +
		"ORDER BY id ASC " +
		"LIMIT $1 " +
		"OFFSET $2;", limitPerPage, page * limitPerPage)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user User

		err := rows.Scan(&user.Id,
			&user.Login,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Image,
			&user.PoolYear,
			&user.PoolMonth,
			&user.Kind,
			&user.Status)

		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	return &users
}

func FindUserById(id int) *User {
	var user User

	row := config.Db().QueryRow("SELECT users.id, " +
																"users.login, " +
																"users.email, " +
																"users.created_at, " +
																"users.updated_at, " +
																"users.image, " +
																"users.pool_year, " +
																"users.pool_month, " +
																"users.kind, " +
																"users.status " +
															"FROM users " +
															"WHERE id = $1;", id)
	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Image,
		&user.PoolYear,
		&user.PoolMonth,
		&user.Kind,
		&user.Status)

	if err != nil {
		log.Fatal(err)
	}

	return &user
}

func FindUserByLogin(login string) *User {
	var user User

	row := config.Db().QueryRow(
		"SELECT users.id, " +
			"users.login, " +
			"users.email, " +
			"users.created_at, " +
			"users.updated_at, " +
			"users.image, " +
			"users.pool_year, " +
			"users.pool_month, " +
			"users.kind, " +
			"users.status " +
		"FROM users " +
		"WHERE login = $1;", login)
	err := row.Scan(&user.Id,
		&user.Login,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Image,
		&user.PoolYear,
		&user.PoolMonth,
		&user.Kind,
		&user.Status)

	if err != nil {
		log.Fatal(err)
	}

	return &user
}