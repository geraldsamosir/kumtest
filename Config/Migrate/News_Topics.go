package Migrate

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type News_Topics struct {
	gorm.Model
	News_id   int  `gorm:"column:news_id;type:int(10)UNSIGNED;"`
	Topics_id  int  `gorm:"column:topics_id;type:int(10)UNSIGNED;"`
}