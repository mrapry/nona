package base

import "time"

//ModelSQL structure define domain for SQL
type ModelSQL struct {
	ID         string     `json:"id"`
	IsDeleted  bool       `json:"is_deleted"`
	CreatedBy  string     `json:"created_by"`
	ModifiedBy string     `json:"modified_by"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at"`
}
