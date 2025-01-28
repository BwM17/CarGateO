package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dtos "mb.com/cargate/Dtos"
	models "mb.com/cargate/Models"
	utils "mb.com/cargate/Utils"
)

func GetLogs(c *gin.Context) {
	var logs []models.Log
	models.DB.Scopes((utils.Paginate(c))).Find(&logs)     

	page := c.GetInt("page")  
	totalPages := c.GetInt("totalPages") 
	pageSize := c.GetInt("pageSize") 
	c.JSON(http.StatusOK, gin.H{
		"page":	page,
		"totalpages": totalPages,
		"pageSize": pageSize,
		"data": logs, 
	})
}  


func Validate(c *gin.Context){ 
	var dto dtos.LogDto 
	var car models.Car 
	now := time.Now()

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.Status(http.StatusBadRequest) 
	} 

	if err := models.DB.Where("numberplate = ?", dto.Numberplate).First(&car).Error; err != nil {    
		log := models.Log{Timewhen: now.Truncate(time.Minute), Numberplate: dto.Numberplate, Allowed: false, Status: fmt.Sprintf("%v is not allowed", car.Numberplate)}
		models.DB.Create(&log)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})  
		return
	}   

	log := models.Log{Timewhen: now.Truncate(time.Minute), Numberplate: dto.Numberplate, Allowed: true, Status: fmt.Sprintf("%v entered successfully", car.Owner)}       
	models.DB.Create(&log)

	c.JSON(http.StatusOK, gin.H{"allowed": true, "paylod": car})
}