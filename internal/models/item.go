package models

type Item struct {
	BaseModel
	RtID        string   `json:"rt_id" gorm:"not null"`
	ItemType    ItemType `json:"item_type" gorm:"not null"`
	Name        string   `json:"name" gorm:"not null"`
	Description string   `json:"Description" gorm:"type:text"`

	Rt Rt `gorm:"foreignKey:RtID" json:"rt"`
}
type Items []Item
