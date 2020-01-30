package models

type Activity struct {
	BaseModel
	Type       int       `json:"type"` // 1 : like, 2 : dislike, 3: comment, 4: purchase
	UserID     *int      `json:"-"`
	User       *User     `json:"user"`
	Content    string    `json:"content"`
	RingtoneID *int      `json:"-"`
	Ringtone   *Ringtone `json:"ringtone"`
}
