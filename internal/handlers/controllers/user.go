package controllers

import (
	"github.com/aliereno/softwareproject-api/internal/handlers/auth"
	"github.com/aliereno/softwareproject-api/internal/orm/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con Controller) UserLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginModel struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := c.BindJSON(&loginModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		user := models.User{}
		db := con.ORM.DB.New()
		er := db.Where("email = ?", loginModel.Email).First(&user).Error
		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
			return
		}
		if err := user.CheckPassword(loginModel.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		token, errr := auth.GetToken(user)
		if errr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": token, "id": user.ID})
	}
}

func (con Controller) UserRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerModel struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		err := c.BindJSON(&registerModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		dbo := &models.User{
			Name:  &registerModel.Name,
			Email: registerModel.Email,
		}
		if err := dbo.SetPassword(registerModel.Password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		tx := con.ORM.DB.New().Begin()
		tx = tx.Create(dbo).First(dbo)
		tx = tx.Commit()
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}

func (con Controller) UsersFetchAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := con.ORM.DB.New()
		var dbRecords []*models.User
		db = db.Find(&dbRecords)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRecords})
	}
}

func (con Controller) UserFetchSetting() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")
		db := con.ORM.DB.New()
		var dbRecord models.User
		err := db.Where("id = ?", id).Preload("Likes").Preload("Dislikes").Preload("Purchases").Find(&dbRecord).Error
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbRecord})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		}
	}
}

func (con Controller) UserUpdateSetting() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "on development..."})
	}
}

func (con Controller) UserPurchaseRingtone() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ringtoneModel struct {
			RingtoneID int `json:"ringtone_id"`
			UserID     int `json:"user_id"`
		}
		err := c.BindJSON(&ringtoneModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}

		var user models.User
		var ringtone models.Ringtone

		tx := con.ORM.DB.New()

		err = tx.Where("id = ?", ringtoneModel.UserID).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		err = tx.Where("id = ?", ringtoneModel.RingtoneID).First(&ringtone).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		tx.Model(user).Association("Purchases").Append(ringtone)
		dboActivity := &models.Activity{
			UserID:     &user.ID,
			Content:    "User " + *user.Name + " purchased Ringtone: " + ringtone.Title,
			RingtoneID: &ringtone.ID,
			Type:       4,
		}
		tx.Create(dboActivity).First(dboActivity)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}
func (con Controller) UserLikeRingtone() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ringtoneModel struct {
			RingtoneID int `json:"ringtone_id"`
			UserID     int `json:"user_id"`
		}
		err := c.BindJSON(&ringtoneModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}

		var user models.User
		var ringtone models.Ringtone

		tx := con.ORM.DB.New()

		err = tx.Where("id = ?", ringtoneModel.UserID).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		err = tx.Where("id = ?", ringtoneModel.RingtoneID).First(&ringtone).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		tx.Model(user).Association("Dislikes").Delete(ringtone)
		tx.Model(user).Association("Likes").Append(ringtone)

		dboActivity := &models.Activity{
			UserID:     &user.ID,
			Content:    "User " + *user.Name + " liked Ringtone: " + ringtone.Title,
			RingtoneID: &ringtone.ID,
			Type:       4,
		}
		tx.Create(dboActivity).First(dboActivity)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}

func (con Controller) UserDislikeRingtone() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ringtoneModel struct {
			RingtoneID int `json:"ringtone_id"`
			UserID     int `json:"user_id"`
		}
		err := c.BindJSON(&ringtoneModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}

		var user models.User
		var ringtone models.Ringtone

		tx := con.ORM.DB.New()

		err = tx.Where("id = ?", ringtoneModel.UserID).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		err = tx.Where("id = ?", ringtoneModel.RingtoneID).First(&ringtone).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}

		tx.Model(user).Association("Likes").Delete(ringtone)
		tx.Model(user).Association("Dislikes").Append(ringtone)
		dboActivity := &models.Activity{
			UserID:     &user.ID,
			Content:    "User " + *user.Name + " disliked Ringtone: " + ringtone.Title,
			RingtoneID: &ringtone.ID,
			Type:       4,
		}
		tx.Create(dboActivity).First(dboActivity)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}
func (con Controller) UserCommentRingtone() gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentModel struct {
			RingtoneID int    `json:"ringtone_id"`
			UserID     int    `json:"user_id"`
			Comment    string `json:"comment"`
		}
		err := c.BindJSON(&commentModel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}

		dbo := &models.Comment{
			UserID:     &commentModel.UserID,
			Comment:    &commentModel.Comment,
			RingtoneID: &commentModel.RingtoneID,
		}
		tx := con.ORM.DB.New().Begin()
		err = tx.Create(dbo).First(dbo).Commit().Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	}
}

func (con Controller) UserFetchRecentActivities() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := con.ORM.DB.New()
		var dbActivities []*models.Activity
		err := db.Table("activities").Order("created_at desc").Limit(10).Find(&dbActivities).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": dbActivities})
	}
}
