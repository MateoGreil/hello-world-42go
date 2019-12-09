package	models

import (
	"github.com/mateogreil/hello-world-42go/config"
	"github.com/mateogreil/hello-world-42go/utils"
	"log"
	"time"
)

type Project	struct {
	Id					int							`json:"id"`
	Name				string					`json:"name"`
	CreatedAt		time.Time				`json:"created_at"`
	UpdatedAt		time.Time				`json:"updated_at"`
	Visible			bool						`json:"visible"`
	Exam				bool						`json:"exam"`
	ParentId		utils.NullInt64	`json:"parent_id"`
	Slug				string					`json:"slug"`
	HasGit			bool						`json:"has_git"`
	HasMark			bool						`json:"has_mark"`
}

type Projects	[]Project

func AllProject(page int) *Projects {
	var limitPerPage = 10
	var projects Projects

	rows, err := config.Db().Query(
		"SELECT projects.id, " +
			"projects.name, " +
			"projects.created_at, " +
			"projects.updated_at, " +
			"projects.visible, " +
			"projects.exam, " +
			"projects.parent_id, " +
			"projects.slug, " +
			"projects.has_git, " +
			"projects.has_mark " +
		"FROM projects " +
		"ORDER BY id ASC " +
		"LIMIT $1 " +
		"OFFSET $2;", limitPerPage, page * limitPerPage)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var project Project

		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.CreatedAt,
			&project.UpdatedAt,
			&project.Visible,
			&project.Exam,
			&project.ParentId,
			&project.Slug,
			&project.HasGit,
			&project.HasMark)

		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, project)
	}

	return &projects
}

func FindProjectById(id int) *Project {
	var project Project

	row := config.Db().QueryRow(
		"SELECT projects.id, " +
			"projects.name, " +
			"projects.created_at, " +
			"projects.updated_at, " +
			"projects.visible, " +
			"projects.exam, " +
			"projects.parent_id, " +
			"projects.slug, " +
			"projects.has_git, " +
			"projects.has_mark " +
		"FROM projects " +
		"WHERE id = $1;", id)

	err := row.Scan(
		&project.Id,
		&project.Name,
		&project.CreatedAt,
		&project.UpdatedAt,
		&project.Visible,
		&project.Exam,
		&project.ParentId,
		&project.Slug,
		&project.HasGit,
		&project.HasMark)

	if err != nil {
		log.Fatal(err)
	}

	return &project
}

