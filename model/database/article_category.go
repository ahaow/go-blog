package database

type ArticleCategory struct {
	Category string `json:"category" grom:"primaryKey"` // 类别
	Number   int    `json:"number"`                     // 统计数量
}
