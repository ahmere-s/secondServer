package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "backend/handlers"
)

func main(){
  r := gin.Default()
  
  //enable CORS for frontend
  r.Use(cors.Default())

  //health check
  r.GET("/health", func(c *gin.Context){
    c.JSON(200, gin.H{"status": "ok"})
  })

  //user routes
  r.GET("/users", handlers.GetUsers)
  r.POST("/users", handlers.CreateUser)

  r.Static("/static", "../frontend") //Serving html at root

  //run backend on port 8080
  r.Run(":8080")
}
