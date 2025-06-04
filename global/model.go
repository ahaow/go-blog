package global

import "time"

type MODEL struct {
	ID       uint      `json:"id" gorm:"primarykey"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"-" gorm:"index"`
}
