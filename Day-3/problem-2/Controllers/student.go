package Controllers

import (
	"Freshers_2021/Day-3/problem-2/Config"
	"Freshers_2021/Day-3/problem-2/Models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func GetUsers(c *gin.Context) {
	student := &[]Models.Student{}
	err := Config.DB.Find(student).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, *student)
	}
}

func CreateUser(c *gin.Context) {
	student := &Models.Student{}
	c.BindJSON(student)
	err := Config.DB.Create(student).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, *student)
	}
}

func GetUserByID(c *gin.Context) {
	student := &Models.Student{}
	id := c.Params.ByName("id")
	err := Config.DB.Where("id = ?", id).First(student).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, *student)
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	student := &Models.Student{}
	err := Config.DB.Where("id = ?", id).First(student).Error
	c.BindJSON(student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		Config.DB.Save(student)
		c.JSON(http.StatusOK, *student)
	}
}

func DeleteUser(c *gin.Context) {
	student := &Models.Student{}
	id := c.Params.ByName("id")
	Config.DB.Where("id = ?", id).Delete(student)
	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}
