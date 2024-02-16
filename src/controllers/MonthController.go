package controllers

import (
	"latihan-gin/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MonthList(c *gin.Context) {
	res := models.MonthList()
	c.JSON(200, gin.H{
		"message": "Successful",
		"data":    res,
	})
}

func MonthShow(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	res := models.MonthDetail(id)
	c.JSON(200, gin.H{
		"message": "Successful",
		"data":    res,
	})
}

func MonthCreate(c *gin.Context) {
	var input models.Month
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.MonthCreate(input.Name, input.Day)
	c.JSON(200, gin.H{
		"message": "created successfully",
	})
}

func MonthUpdate(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input models.Month
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.MonthUpdate(id, input.Name, input.Day)
	c.JSON(200, gin.H{
		"message": "Updated successfully",
	})
}

func MonthDestroy(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	models.MonthDelete(id)
	c.JSON(200, gin.H{
		"message": "Deleted successfully",
	})
}
