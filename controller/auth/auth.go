package auth

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"example.com/greetings/model"
	"example.com/greetings/orm"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	UserID   uint   `json:"user_id"`
	Username string ` json:"username" binding:"required"`
	Password string ` json:"password" binding:"required"`
	Fullname string ` json:"fullname" binding:"required"`
	Avatar   string ` json:"avatar" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if username already exists
	var userExist model.User
	if err := orm.DB.Where("username = ?", json.Username).First(&userExist).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	// create user
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.User{
		UserID:   uint(uuid[0]),
		Username: json.Username,
		Password: string(hashedPassword),
		Fullname: json.Fullname,
		Avatar:   json.Avatar,
	}
	orm.DB.Create(&user)
	if user.ID != 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User created successfully", "data": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	}
}

// Login is a function that handles the login request
type LoginBody struct {
	Username string ` json:"username" binding:"required"`
	Password string ` json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := orm.DB.Where("username = ?", json.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	} else {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["id"] = user.ID
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KE")))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Login success", "status": "success", "token": t, "user": user})
	}

}
