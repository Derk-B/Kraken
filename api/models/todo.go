package models

import "github.com/uptrace/bun"

type Todo struct {
	bun.BaseModel `bun:"table:todos`

	Id        int64 `bun:",pk,autoincrement"`
	UserId    int64
	Title     string
	Completed bool
}
