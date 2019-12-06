package	models

import (
	"github.com/synbioz/go_api/config"
	"log"
	"time"
)

type ProjectsUser			struct {
	Id				int				`json:"id"`
	UserId		int				`json:"user_id"`
	ProjectId	int				`json:"project_id"`
	MarkedAt	time.time	`json:"marked_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	CreatedAt	time.Time	`json:"created_at"`
}

type ProjectsUsers		[]ProjectsUser