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
	if err := orm.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}
