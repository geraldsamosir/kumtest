package Migrate

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Topics struct {
	gorm.Model
	Topic_name string `gorm:"column:topic_name;not null" `
	Parent_id int  `gorm:"column:parent_id;type:int(10)UNSIGNED;DEFAULT:Null"`
	News []News `gorm:"many2many:news_topics;save_association:false"`
}