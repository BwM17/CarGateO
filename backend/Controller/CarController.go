package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtos "mb.com/cargate/Dtos"
	models "mb.com/cargate/Models"
	utils "mb.com/cargate/Utils"
)


func GetCars(c *gin.Context) {
	var cars []models.Car
	models.DB.Scopes((utils.Paginate(c))).Find(&cars)    

	page := c.GetInt("page")  
	totalPages := c.GetInt("totalPages") 
	pageSize := c.GetInt("pageSize") 
	c.JSON(http.StatusOK, gin.H{
		"page":	page,
		"totalpages": totalPages,
		"pageSize": pageSize,
		"data": cars, 
	})
} 

func GetCar(c *gin.Context) {
}  


func PostCar(c *gin.Context) {
	var dto dtos.CarDto 

	if err := c.ShouldBindJSON(&dto); err != nil { 
		c.Status(http.StatusBadRequest)   
	}  

	car := models.Car{Model: dto.Model, Numberplate: dto.Numberplate, Owner: dto.Owner, Type: dto.Type, Status: dto.Status, LeasedAt: dto.LeasedAt, ReturnAt: dto.ReturnAt} 
	models.DB.Create(&car) 

	c.JSON(http.StatusOK, car)  
}

func PatchCar(c *gin.Context) { 
	var car models.Car
	var dto dtos.CarDto  

	if err := models.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Record not found!"})  
		return
	}

	if err := c.ShouldBindJSON(&dto); err != nil { 
		c.Status(http.StatusBadRequest)   
	}   

	models.DB.Model(&car).Updates(&dto)  
	c.JSON(http.StatusOK, dto) 
} 

func DeleteCar(c *gin.Context) {
	var car models.Car
	if err := models.DB.Where("id = ?", c.Param("id")).First(&car).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"}) 
		return
	} 

	models.DB.Delete(&car)  

	c.JSON(http.StatusOK, gin.H{"data": true}) 
}