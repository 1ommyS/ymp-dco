package domain

import "github.com/google/uuid"

type Group struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
}

func NewGroup(id uuid.UUID, title string) Group {
	return Group{
		ID:    id,
		Title: title,
	}
}
