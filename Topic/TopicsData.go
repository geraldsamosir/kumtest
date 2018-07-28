package Topic

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/geraldsamosir/kumtest/Config/Migrate"	
)

type TopicsData struct {
	Migrate.Topics
}

func (t *TopicsData) TableName() string {
	return "topics"
}

func GetTopics(db *gorm.DB, ids int, t *TopicsData) (err error) {
	if err := db.Where("id = ?", ids).First(&t).Error; err != nil {
		return err
	}
	return nil
}


func GetAllTopics(db *gorm.DB, filter *TopicsData ,t *[]TopicsData) (err error) {
	if err = db.Preload("News").Where(filter).Order("topic_name desc").Find(t).Error; err != nil  {
		return err
	}
	return nil
}



func GetAllTopicsDeleted(db *gorm.DB, filter *TopicsData  ,t *[]TopicsData) (err error) {
	if err = db.Preload("News").Unscoped().Where(filter).Not("deleted_at = ?", "Null").Find(t).Error; err != nil  {
		return err
	}
	return nil
}

func InsertTopics(db *gorm.DB, t *TopicsData) (err error) {
	if err = db.Create(t).Error; err != nil {
		return err
	}
	return nil
}



func UpdateTopics(db *gorm.DB, t *TopicsData) (err error) {
	if err = db.Save(t).Error; err != nil {
		return err
	}
	return nil
}

func DeletedTopics(db *gorm.DB, t *TopicsData) (err error) {
	if err = db.Delete(t).Error; err != nil {
		return err
	}
	return nil
}