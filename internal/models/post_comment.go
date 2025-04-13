package models

type PostComment struct {
	BaseModel
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
	Comment string `json:"comment"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
}
