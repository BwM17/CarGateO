package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dtos "mb.com/cargate/Dtos"
	models "mb.com/cargate/Models"
	utils "mb.com/cargate/Utils"
)


func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Scopes((utils.Paginate(c))).Find(&users)     

	page := c.GetInt("page")  
	totalPages := c.GetInt("totalPages") 
	pageSize := c.GetInt("pageSize") 
	c.JSON(http.StatusOK, gin.H{
		"page":	page,
		"totalpages": totalPages,
		"pageSize": pageSize,
		"data": users, 
	})
} 

func GetUser(c *gin.Context) {
}  


func PostUser(c *gin.Context) {
	var dto dtos.UserDto 

	if err := c.ShouldBindJSON(&dto); err != nil { 
		c.Status(http.StatusBadRequest)   
	}   

	now := time.Now()

	user := models.User{Sirname: dto.Sirname, Lastname: dto.Lastname, Email: dto.Email, Role: dto.Role, Userage: now.Truncate(time.Minute)}  
	models.DB.Create(&user)  

	c.JSON(http.StatusOK, user)   
}

func PatchUser(c *gin.Context) { 
	var user models.User
	var dto dtos.UserDto  

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error":"Record not found!"})  
		return
	}

	if err := c.ShouldBindJSON(&dto); err != nil { 
		c.Status(http.StatusBadRequest)   
	}   

	models.DB.Model(&user).Updates(&dto)  
	c.JSON(http.StatusOK, dto) 
} 

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"}) 
		return
	} 

	models.DB.Delete(&user)   

	c.JSON(http.StatusOK, gin.H{"data": true}) 
}