package Migrate

import (
//	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&News_Topics{},&News{},&Topics{})
	db.Model(&Topics{}).AddForeignKey("parent_id","topics(id)", "RESTRICT", "RESTRICT")
   db.Model(&News_Topics{}).AddForeignKey("news_id","news(id)", "RESTRICT", "RESTRICT").AddForeignKey("topics_id","topics(id)", "RESTRICT", "RESTRICT")

	//topicTableName := db.NewScope(&Topics{}).GetModelStruct().TableName(db)

	//db.Model(&Topics{}).AddForeignKey("parent_id", topicTableName + "(id)", "RESTRICT", "RESTRICT")
	// hasUser := db.HasTable(&News{})
	// fmt.Println("News is ", hasUser)
	// if !hasUser {
	// 	db.CreateTable(&News{})
	// }

	return db
}
