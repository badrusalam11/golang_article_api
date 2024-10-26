package models

import "time"

type Article struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"type:varchar(200);not null"`
	Content     string    `gorm:"type:text;not null"`
	Category    string    `gorm:"type:varchar(100);not null"`
	Status      string    `gorm:"type:varchar(100);not null"` // "publish", "draft", "thrash"
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

type ArticleResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category    string    `json:"category"`
	Status      string    `json:"status"` // "publish", "draft", "thrash"
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

// ToResponse converts an Article to an ArticleResponse
func (a *Article) ToResponse() ArticleResponse {
	return ArticleResponse{
		ID:          a.ID,
		Title:       a.Title,
		Content:     a.Content,
		Category:    a.Category,
		Status:      a.Status,
		CreatedDate: a.CreatedDate,
		UpdatedDate: a.UpdatedDate,
	}
}

func (Article) TableName() string {
	return "posts"
}
