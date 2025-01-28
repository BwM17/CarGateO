package dtos

type UserDto struct {
	Sirname 	string		`json:"sirname"`
	Lastname 	string 		`json:"lastname"`  
	Email 		string 		`json:"email"` 
	Role 		string 		`json:"role"`
}