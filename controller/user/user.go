package user

import (
	"net/http"

	"example.com/greetings/model"
	"example.com/greetings/orm"
	"github.com/gin-gonic/gin"
)

// GetUser is a function that handles the get user request
func GetUsers(c *gin.Context) {
	var users []model.User
	orm.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": users})

}

// GetUser is a function that handles the get user request
func GetUser(c *gin.Context) {
	var user model.User
	if err := orm.DB.Find(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	if user.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User found successfully", "data": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "User not found!"})
	}
}
