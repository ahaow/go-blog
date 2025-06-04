package database

type ArticleTag struct {
	Tag    string `json:"tag" grom:"primaryKey"` // 标签
	Number int    `json:"number"`                // 统计数量
}
