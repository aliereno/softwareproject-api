package controllers

import (
	"github.com/aliereno/softwareproject-api/internal/orm/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con Controller) CategoriesFetchAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := con.ORM.DB.New()
		var dbRecords []*models.Category
		db = db.Preload("Ringtones").Find(&dbRecords)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRecords})
	}
}

func (con Controller) CategoriesFetchDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		db := con.ORM.DB.New()
		var dbRecord models.Category
		err := db.Preload("Ringtones").Where("id = ?", id).Find(&dbRecord).Error
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRecord})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": ""})
		}
	}
}
