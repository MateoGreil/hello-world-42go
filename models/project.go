package	models

import (
	"github.com/synbioz/go_api/config"
	"log"
	"time"
)

type Project 	struct {
	Id				int				`json:"id"`
	Name			string		`json:"name"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type Projects	[]Project