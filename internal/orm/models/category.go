package models

type Category struct {
	BaseModel
	Title     *string     `json:"title"`
	Ringtones []*Ringtone `json:"ringtones"`
}
