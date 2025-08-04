package dto

type GetAllSegmentsDbResult struct {
	Title     string `db:"title"`
	P         uint32 `db:"p"`
	GroupName string `db:"group_name"`
	Response  string `db:"response"`
}
