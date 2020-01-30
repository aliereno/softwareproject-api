package models

type Comment struct {
	BaseModel
	UserID     *int      `json:"-"`
	User       *User     `json:"user"`
	Comment    *string   `json:"comment"`
	RingtoneID *int      `json:"-"`
	Ringtone   *Ringtone `json:"ringtone"`
}
