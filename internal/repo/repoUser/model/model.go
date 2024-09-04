package repo

import (
	"database/sql"
	"time"
)

type RepoUser struct {
	Id        int64
	Name      string
	Password  string
	Role      int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
