package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "backend/models"
)

func GetUsers(c *gin.Context){
    c.JSON(http.StatusOK, models.Users)
}

func CreateUser(c *gin.Context){
    var newUser models.User

    if err := c.BindJSON(&newUser); err != nil {
	    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	    return
    }

    newUser.ID = len(models.Users) + 1
    models.Users = append(models.Users, newUser)

    c.JSON(http.StatusCreated, newUser)
}
