
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

func (Article) TableName() string {
    return "posts"
}
