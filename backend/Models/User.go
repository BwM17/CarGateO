package models

import "time"

type User struct {
	ID 			uint 		`json:"id" gorm:"primary key"` 
	Sirname 	string		`json:"sirname"`
	Lastname 	string 		`json:"lastname"` 
	Email		string	 	`json:"email"`
	Role 		string 		`json:"role"`  
	Userage		time.Time	`json:"userage"`   
}