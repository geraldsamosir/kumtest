package News_topics

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/geraldsamosir/kumtest/Config/Migrate"	
)

type NewsTopicData struct {
	Migrate.News_Topics
}

func (t *NewsTopicData) TableName() string {
	return "news_topics"
}

func GetTopics(db *gorm.DB, ids int, t *NewsTopicData) (err error) {
	if err := db.Where("id = ?", ids).First(&t).Error; err != nil {
		return err
	}
	return nil
}


func GetAllTopics(db *gorm.DB, filter *NewsTopicData ,t *[]NewsTopicData) (err error) {
	if err = db.Preload("News").Where(filter).Find(t).Error; err != nil  {
		return err
	}
	return nil
}



func GetAllTopicsDeleted(db *gorm.DB, filter *NewsTopicData  ,t *[]NewsTopicData) (err error) {
	if err = db.Preload("News").Unscoped().Where(filter).Not("deleted_at = ?", "Null").Find(t).Error; err != nil  {
		return err
	}
	return nil
}

func InsertTopics(db *gorm.DB, t *NewsTopicData) (err error) {
	if err = db.Create(t).Error; err != nil {
		return err
	}
	return nil
}



func UpdateTopics(db *gorm.DB, t *NewsTopicData) (err error) {
	if err = db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func DeletedTopics(db *gorm.DB, t *NewsTopicData) (err error) {
	if err = db.Model(&t).Association("Languages").Clear().Error; err != nil {
		return err
	}
	return nil
}