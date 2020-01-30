package controllers

import (
	"github.com/aliereno/softwareproject-api/internal/orm/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func (con Controller) RingtonesFetchAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := con.ORM.DB.New()
		var dbRecords []*models.Ringtone
		db = db.Preload("Category").Preload("Likes").Preload("Dislikes").Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User")
		}).Find(&dbRecords)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRecords})
	}
}

func (con Controller) RingtoneFetchDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		db := con.ORM.DB.New()
		var dbRecord models.Ringtone
		err := db.Preload("Category").Preload("Likes").Preload("Dislikes").Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User")
		}).Where("id = ?", id).Find(&dbRecord).Error
		if err == nil {
			dbRecord.Hit += 1
			db.Save(&dbRecord)
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRecord})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "data": ""})
		}
	}
}

//TODO: ikinci sprintte son aksiyonlar tarzı akış; şu kullanıcı bunu beğendi, satın aldı.
func (con Controller) RingtonesFetchFeatued() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := con.ORM.DB.New()
		var dbRingtones []*models.Ringtone
		err := db.Preload("Category").Preload("Likes").Preload("Dislikes").Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User")
		}).Order("hit desc").Limit(3).Find(&dbRingtones).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRingtones})
	}
}
