package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users`

	Id       int64 `bun:",pk,autoincrement"`
	Username string
	Email    string
	Password string
}
