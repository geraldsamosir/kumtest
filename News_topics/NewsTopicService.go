package News_topics  

import(
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/schema"
	
)

var decoder  = schema.NewDecoder()



type NewsTopicService struct {
	DB *gorm.DB
}

func (T * NewsTopicService) setGorm(db *gorm.DB){
	T.DB =  db
}

func ( N * NewsTopicService ) Get(res http.ResponseWriter, req *http.Request){
	var  _topic NewsTopicData
   _topics := []NewsTopicData{}
   var deleted_At = req.URL.Query().Get("deleted_at")
   errParseURl := decoder.Decode(&_topic, req.URL.Query())
   if errParseURl != nil  {
	   res.Header().Set("Content-Type", "application/json")
	   res.WriteHeader(500) 
	   json.NewEncoder(res).Encode(errParseURl.Error())
	   return
   }
   errDB:= GetAllTopics(N.DB,&_topic ,&_topics)
   if deleted_At == "true" {
	   errDB = GetAllTopicsDeleted(N.DB,&_topic ,&_topics)		
   }



   if errDB != nil {
	   res.Header().Set("Content-Type", "application/json")
	   res.WriteHeader(500)
	   json.NewEncoder(res).Encode(errDB.Error())
	   return
   }
   
   res.Header().Set("Content-Type", "application/json")
   res.WriteHeader(200)
   json.NewEncoder(res).Encode(_topics)
   return
}

func ( T * NewsTopicService ) GetDeletedTopcis(res http.ResponseWriter, req *http.Request){
    var  _topic NewsTopicData
   _topics := []NewsTopicData{}
   errParseURl := decoder.Decode(&_topic, req.URL.Query())
   if errParseURl != nil  {
	   res.Header().Set("Content-Type", "application/json")
	   res.WriteHeader(500) 
	   json.NewEncoder(res).Encode(errParseURl.Error())
	   return
   }
   
   errDB := GetAllTopicsDeleted(T.DB,&_topic ,&_topics)		
   
   if errDB != nil {
	   res.Header().Set("Content-Type", "application/json")
	   res.WriteHeader(500)
	   json.NewEncoder(res).Encode(errDB.Error())
	   return
   }
   
   res.Header().Set("Content-Type", "application/json")
   res.WriteHeader(200)
   json.NewEncoder(res).Encode(_topics)
   return
}


func (T *NewsTopicService)Create(res http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)
	var _topics NewsTopicData
	//var _Topics []Migrate.Topics
	err := decoder.Decode(&_topics) 
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	defer req.Body.Close()
	err = InsertTopics(T.DB, &_topics)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500) 
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(201) 
	json.NewEncoder(res).Encode(_topics)
} 


func (T *NewsTopicService)Update(res http.ResponseWriter, req *http.Request){
	
	var _topics NewsTopicData
	var id = req.URL.Query().Get("id")
	_id, err := strconv.Atoi(id)
	err= GetTopics(T.DB, _id , &_topics)
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&_topics) 
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	defer req.Body.Close()
	err = UpdateTopics(T.DB, &_topics)
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	 }
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200) 
	json.NewEncoder(res).Encode(_topics)
}

func (T *NewsTopicService)Delete(res http.ResponseWriter, req *http.Request){
	var _topics NewsTopicData
	var id = req.URL.Query().Get("id")
	_id, err := strconv.Atoi(id)
	err = GetTopics(T.DB, _id , &_topics)
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	err = DeletedTopics(T.DB , &_topics)
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	 }
	 res.Header().Set("Content-Type", "application/json")
	 res.WriteHeader(200) 
	 json.NewEncoder(res).Encode(_topics)
}