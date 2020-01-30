package models

type Ticket struct {
	BaseModel
	UserID  *int    `json:"-"`
	User    *User   `json:"user"`
	Content *string `json:"content"`
}
