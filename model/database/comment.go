package database

import (
	"github.com/gofrs/uuid"
	"go-blog/global"
)

// Comment 评论表

type Comment struct {
	global.MODEL
	ArticleID string    `json:"article_id"` // 文章id
	PID       *uint     `json:"p_id"`       // 父评论 id
	PComment  *Comment  `json:"-" gorm:"foreignKey:PID"`
	Children  []Comment `json:"children" grom:"foreignKey:PID"`                  // 子评论
	UserUUID  uuid.UUID `json:"user_uuid" gorm:"type:char(36)"`                  // 用户 uuid
	User      User      `json:"user" gorm:"foreignKey:UserUUID;references:UUID"` // 关联的用户
	Content   string    `json:"content"`                                         // 内容
}
