package models

import "time"

type Log struct {
	ID 				uint		`json:"id" gorm:"primary key"`  
	Timewhen		time.Time	`json:"timewhen"`
	Numberplate		string		`json:"numberplate"`    
	Allowed			bool		`json:"allowed"`
	Status			string		`json:"status"`
}