package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/models"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"net/http"
	"github.com/itsbalamurali/heyasha/config"
	"github.com/jinzhu/gorm"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest,gin.H{"code":400,"message":err.Error()})
	}else {
		db := c.MustGet("mysql").(*gorm.DB)
		passbyte := []byte(user.Password)
		md5pass := md5.Sum(passbyte)
		usr := &models.User{
			Username: user.Username,
			Email:user.Email,
			Password: hex.EncodeToString(md5pass[:]),
		}
		db.Save(&usr)
		// Create the token
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
		// Set some claims
		token.Claims["ID"] = usr.ID
		//token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(config.AppSecret))
		if err != nil {
			c.JSON(500, gin.H{"code":500, "message": "Could not generate token"})
		}
		session := &models.Session{
			UserID: usr.ID,
			SessionToken: tokenString,
			IPAddress:c.Request.RemoteAddr,
		}
		db.Save(&session)
		c.Header("Location", "https://api.heyasha.com/v1/users/"+strconv.Itoa(int(usr.ID)))
		c.JSON(http.StatusCreated, gin.H{"objectId":session.ID,"createdAt":session.CreatedAt,"sessionToken": tokenString})
	}
}

func LoginUser(c *gin.Context) {
	login := Login{}
	if err := c.BindJSON(&login); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
	}else {
		db := c.MustGet("mysql").(*gorm.DB)
		user := &models.User{}
		db.Where("email = ?", login.User).First(&user)
		if user.ID == 0 {
			c.JSON(200, gin.H{"success":false,"message":"Invalid User Account" })
			c.Abort()
		}
		passbyte := []byte(login.Password)
		md5password := md5.Sum(passbyte)
		md5pass := hex.EncodeToString(md5password[:])
		//Password is okay
		if md5pass == user.Password {
			// Create the token
			token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
			// Set some claims
			token.Claims["ID"] = user.ID
			//token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
			// Sign and get the complete encoded token as a string
			tokenString, err := token.SignedString([]byte(config.AppSecret))
			if err != nil {
				c.JSON(500, gin.H{"code":500,"message": "Could not generate token"})
				c.Abort()
			}
			session := &models.Session{
				UserID: user.ID,
				SessionToken:user.Pid,
				IPAddress:c.Request.RemoteAddr,
			}
			db.Create(&session)
			c.JSON(200, gin.H{"success":true,"token": tokenString})
		}else {
			c.JSON(200, gin.H{"success":false,"message":"Invalid Password" })
		}
		}
}

func GetUserDetails(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
	}else {
		c.JSON(http.StatusFound,user)
	}
}

func UpdateUserDetails(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
	}else {
		c.JSON(http.StatusOK,user)
	}
}

func DeleteUser(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
	}else {
		c.JSON(http.StatusFound,user)
	}
}

func ResetPassword(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
	}else {
		c.JSON(http.StatusFound,user)
	}
}
