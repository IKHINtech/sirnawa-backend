package models

type Post struct {
	BaseModel
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
