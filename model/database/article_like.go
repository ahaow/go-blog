package database

import "go-blog/global"

// 文章收藏表

type ArticleLike struct {
	global.MODEL
	ArticleID string `json:"article_id"` // 文章 ID
	UserId    uint   `json:"user_id"`    // 用户 ID
	User      User   `json:"-" gorm:"foreignKey: UserID"`
}
