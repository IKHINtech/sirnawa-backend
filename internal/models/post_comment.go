package models

type PostComment struct {
	BaseModel
	PostID  string `gorm:"not null"  json:"post_id"`
	UserID  string `gorm:"not null" json:"user_id"`
	Comment string `gorm:"not null" json:"comment"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	Post    Post   `gorm:"foreignKey:PostID" json:"post"`
}
