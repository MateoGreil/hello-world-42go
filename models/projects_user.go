package models

import (
	"github.com/mateogreil/hello-world-42go/config"
	// "github.com/mateogreil/hello-world-42go/utils"
	"log"
	"time"
)

type ProjectsUser struct {
	Id					int				`json:"id"`
	ProjectId		int				`json:"project_id"`
	UserId			int				`json:"user_id"`
	CreatedAt		time.Time `json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
	Occurrence	int				`json:"occurrence"`
	FinalMark		int				`json:"final_mark"`
	RetriableAt	time.Time	`json:"retriable_at"`
	MarkedAt		time.Time	`json:"marked_at"`
}

type ProjectsUsers	[]ProjectsUser

func	FindProjectsUserById(id int) *ProjectsUser {
	var projects_user ProjectsUser

	row := config.Db().QueryRow(
		"SELECT projects_users.id, " +
			"projects_users.project_id, " +
			"projects_users.user_id, " +
			"projects_users.created_at, " +
			"projects_users.updated_at, " +
			"projects_users.occurrence, " +
			"projects_users.final_mark, " +
			"projects_users.retriable_at, " +
			"projects_users.marked_at " +
		"FROM projects_users " +
		"WHERE id = $1", id)
	
	err := row.Scan(
		&projects_user.Id,
		&projects_user.ProjectId,
		&projects_user.UserId,
		&projects_user.CreatedAt,
		&projects_user.UpdatedAt,
		&projects_user.Occurrence,
		&projects_user.FinalMark,
		&projects_user.RetriableAt,
		&projects_user.MarkedAt)
	
	if err != nil {
		log.Fatal(err)
	}

	return &projects_user
}

func	FindProjectsUserByUserId(user_id int) *ProjectsUser {
	var projects_user ProjectsUser

	row := config.Db().QueryRow(
		"SELECT projects_users.id, " +
			"projects_users.project_id, " +
			"projects_users.user_id, " +
			"projects_users.created_at, " +
			"projects_users.updated_at, " +
			"projects_users.occurrence, " +
			"projects_users.final_mark, " +
			"projects_users.retriable_at, " +
			"projects_users.marked_at " +
		"FROM projects_users " +
		"WHERE user_id = $1", user_id)
	
	err := row.Scan(
		&projects_user.Id,
		&projects_user.ProjectId,
		&projects_user.UserId,
		&projects_user.CreatedAt,
		&projects_user.UpdatedAt,
		&projects_user.Occurrence,
		&projects_user.FinalMark,
		&projects_user.RetriableAt,
		&projects_user.MarkedAt)
	
	if err != nil {
		log.Fatal(err)
	}

	return &projects_user
}