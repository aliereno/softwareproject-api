package models

type Ringtone struct {
	BaseModel
	CategoryID *int       `json:"-"`
	Category   *Category  `json:"category"`
	Price      int        `json:"price"`
	FileName   string     `json:"file_name"`
	Title      string     `json:"title"`
	Content    *string    `json:"content"`
	Likes      []*User    `json:"likes" gorm:"many2many:user_ringtone_likes;association_jointable_foreignkey:user_id"`
	Dislikes   []*User    `json:"dislikes" gorm:"many2many:user_ringtone_dislikes;association_jointable_foreignkey:user_id"`
	Comments   []*Comment `json:"comments" gorm:"foreignkey:RingtoneID"`
	Hit        int        `json:"hit" gorm:"default:0"`
}
