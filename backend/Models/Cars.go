package models

type Car struct {
	ID				uint		`json:"id" gorm:"primary key"` 
	Model   		string		`json:"model"`  
	Numberplate 	string 		`json:"numberplate" gorm:"requiered"`    
	Owner			string		`json:"owner"`  
	Type			string 		`json:"type"`
	Status			string		`json:"status"`
	LeasedAt		string		`json:"lease"` 
	ReturnAt		string		`json:"return"`
}