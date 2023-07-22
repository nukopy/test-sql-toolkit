// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"time"
)

type Author struct {
	ID        int32
	Name      string
	CreatedAt time.Time
}

type Book struct {
	ID        int32
	Title     string
	Price     int32
	AuthorID  int32
	CreatedAt time.Time
}