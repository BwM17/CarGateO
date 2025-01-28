package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controller "mb.com/cargate/Controller"
	models "mb.com/cargate/Models"
)


func main() { 
		r := gin.Default()   

	//migrate
	models.ConnectDatabase() 

	//cors config
	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true
    config.MaxAge = 12 * time.Hour

    r.Use(cors.New(config)) 

	//Car Endpoint
	r.GET("/api/car", controller.GetCars) 
	r.POST("/api/car", controller.PostCar) 
	r.PATCH("/api/car/:id", controller.PatchCar) 
	r.DELETE("/api/car/:id", controller.DeleteCar) 

	//User Endpoint
	r.GET("/api/user", controller.GetUsers) 
	r.POST("/api/user", controller.PostUser) 
	r.PATCH("/api/user/:id", controller.PatchUser)  
	r.DELETE("/api/user/:id", controller.DeleteUser)

	//Validation / Log Endpint   
	r.GET("/api/logs", controller.GetLogs)
	r.POST("/api/valid", controller.Validate)  

	err := r.Run(":3030")
        if err != nil { 
				log.Fatalf("Error%v\n",err.Error())
                return
        }
}