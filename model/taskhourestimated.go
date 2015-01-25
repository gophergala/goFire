package model

import "time"

type TaskHourEstimated struct {
	TaskHourEstimatedID	int			`db:"TaskHourEstimatedID"`
	TaskID 				int			`db:"TaskID"`	
	ProjectedHours  	decimal 	`db:"ProjectedHours"`
	RemainingHours 		decimal		`db:"RemainingHours"`
	PersonID 			int   		`db:"PersonID"`
	CreateDate   		time.time 	`db:"CreateDate"`
	CreatedPersonID 	int   		`db:"CreatedPersonID"`
}
