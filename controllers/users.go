package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/models"
	"net/http"
)

func CreateUser(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
	}else {
		c.JSON(http.StatusCreated,user)
	}


}

func LoginUser(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(&user); err != nil {
		c.Error(err)
	}else {
		c.JSON(http.StatusOK,user)
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
