package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Email        string  `json:"email" gorm:"not null;unique_index:idx_email"`
	PasswordHash string  `json:"-" gorm:"column:password;not null"`
	Name         *string `json:"name"`
	Role         *int    `json:"-" gorm:"default:0"`
	//Wishes        		[]*Ringtone     `gorm:"many2many:user_ringtone_wishes;association_jointable_foreignkey:ringtone_id"`
	Likes     []*Ringtone `json:"likes" gorm:"many2many:user_ringtone_likes;association_jointable_foreignkey:ringtone_id"`
	Dislikes  []*Ringtone `json:"dislikes" gorm:"many2many:user_ringtone_dislikes;association_jointable_foreignkey:ringtone_id"`
	Comments  []*Comment  `json:"comments" gorm:"foreignkey:UserID"`
	Purchases []*Ringtone `json:"purchases" gorm:"many2many:user_ringtone_purchases;association_jointable_foreignkey:ringtone_id"`
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty!")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

// Database will only save the hashed string, you should check it by util function.
// 	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *User) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
