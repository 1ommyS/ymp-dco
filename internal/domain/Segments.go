package domain

import "github.com/google/uuid"

type Segments struct {
	ID       uuid.UUID `db:"id"`
	Title    string    `db:"title"`
	GroupId  uuid.UUID `db:"group_id"`
	P        uint32    `db:"p"`
	Response string    `db:"response"`
}

func NewSegments(id uuid.UUID, title string) Group {
	return Group{
		ID:    id,
		Title: title,
	}
}
