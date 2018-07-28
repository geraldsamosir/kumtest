package News

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/geraldsamosir/kumtest/Config/Migrate"
	"github.com/geraldsamosir/kumtest/Topic"	
)


type NewsData struct {
	//gorm.Model
	Migrate.News
}

func (n *NewsData) TableName() string {
	return "news"
  }

func GetAllNews(db *gorm.DB, filter *NewsData  ,n *[]NewsData) (err error) {
	if err = db.Preload("Topics").Where(filter).Order("title desc").Find(n).Error; err != nil  {
		return err
	}
	return nil
}

func GetAllNewsBYTopics(db *gorm.DB , filter * Topic.TopicsData  ,n *[]NewsData) (err error) {
	var _news   NewsData
	if err = db.Model(&filter).Related(&_news,"news").Find(&n).Error; err != nil  {
		return err
	}
	return nil
}




func GetNews(db *gorm.DB, ids int, n *NewsData) (err error) {
	if err := db.Where("id = ?", ids).First(&n).Error; err != nil {
		return err
	}
	return nil
}


func GetAllNewsDeleted(db *gorm.DB, filter *NewsData  ,n *[]NewsData) (err error) {
	if err = db.Preload("Topics").Unscoped().Where(filter).Not("deleted_at = ?", "Null").Find(n).Error; err != nil  {
		return err
	}
	return nil
}


func InsertNews(db *gorm.DB, n *NewsData) (err error) {
	if err = db.Create(n).Error; err != nil {
		return err
	}
	return nil
}



func UpdateNews(db *gorm.DB,  _id int,  n *NewsData) (err error) {
	if err = db.Save(n).Error; err != nil {
		return err
	}
	return nil
}

func DeletedNews(db *gorm.DB, n *NewsData) (err error) {
	if err = db.Delete(n).Error; err != nil {
		return err
	}
	return nil
}