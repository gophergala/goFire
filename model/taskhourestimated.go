package model

import "time"

type TaskHourEstimated struct {
	TaskHourEstimatedID int       `db:"TaskHourEstimatedID"`
	TaskID              int       `db:"TaskID"`
	ProjectedHours      int       `db:"ProjectedHours"`
	RemainingHours      int       `db:"RemainingHours"`
	PersonID            int       `db:"PersonID"`
	CreateDate          time.Time `db:"CreateDate"`
	CreatedPersonID     int       `db:"CreatedPersonID"`
}
