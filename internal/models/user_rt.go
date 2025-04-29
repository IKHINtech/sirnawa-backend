package models

type UserRT struct {
	BaseModel
	UserID string `json:"user_id" gorm:"not null;index:idx_user_rt_role,unique"`
	RtID   string `json:"rt_id" gorm:"not null;index:idx_user_rt_role,unique"`
	Role   Role   `json:"role" gorm:"not null;index:idx_user_rt_role,unique"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Rt     Rt     `json:"rt" gorm:"foreignKey:RtID"`
}

type UserRTs []UserRT
