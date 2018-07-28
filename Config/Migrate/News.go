package Migrate

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type News struct {
	gorm.Model
	Title string `gorm:"column:title;not null"`
	Body string `gorm:"column:body;not null"`
	Banner string  `gorm:"column:banner;not null"`
	DatePublish time.Time `gorm:"column:datePublish;not null;DEFAULT:current_timestamp"`
	Is_draf bool `gorm:"column:is_draf;not null;default:true"`
	Is_publish bool `gorm:"column:is_publish;not null;default:false"`
	Topics []Topics `gorm:"many2many:news_topics;save_association:false"`
}