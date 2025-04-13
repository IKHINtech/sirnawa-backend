package models

type PostComment struct {
	BaseModel
	PostID  string `json:"post_id"`
	UserID  string `json:"user_id"`
	Comment string `json:"comment"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
}
