package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Role      int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
