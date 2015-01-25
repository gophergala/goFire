package model

import "time"

type Sprint struct {
	SprintID      	int
	SprintTitle   	string `sql:"type:varchar(45);"`
	Description    	string `sql:"type:varchar(45);"`
	CompanyID       int
	StartDate       time.Time
	EndDate 		time.Time
}
