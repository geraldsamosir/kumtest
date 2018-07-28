package  News

import(
	"os"
	"fmt"
	"io"
	"time"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/schema"
	"github.com/geraldsamosir/kumtest/Topic"
	"github.com/geraldsamosir/kumtest/Config"
)

var decoder  = schema.NewDecoder()

type NewsService struct {
	DB *gorm.DB
}


func (N * NewsService) setGorm(db *gorm.DB){
	N.DB =  db
}

func ( N * NewsService ) Get(res http.ResponseWriter, req *http.Request){
	 var  _new NewsData
	_news := []NewsData{}
	errParseURl := decoder.Decode(&_new, req.URL.Query())
	if errParseURl != nil  {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500) 
		json.NewEncoder(res).Encode(errParseURl.Error())
		return
	}
	errDB:= GetAllNews(N.DB,&_new ,&_news)


	if errDB != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(errDB.Error())
		return
	}
	
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(_news)
	return
}

func (N * NewsService) GetbyTopics(res http.ResponseWriter, req *http.Request){
	var  _topics Topic.TopicsData
	_news := []NewsData{}
	errParseURl := decoder.Decode(&_topics, req.URL.Query())
	if errParseURl != nil  {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500) 
		json.NewEncoder(res).Encode(errParseURl.Error())
		return
	}
	errDB:= GetAllNewsBYTopics(N.DB,&_topics ,&_news)


	if errDB != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(errDB.Error())
		return
	}
	
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(_news)
	return
}


func ( N * NewsService ) GetDeletedNews(res http.ResponseWriter, req *http.Request){
	var  _new NewsData
   _news := []NewsData{}
   errParseURl := decoder.Decode(&_new, req.URL.Query())
   if errParseURl != nil  {
	   res.Header().Set("Content-Type", "application/json")
	   res.WriteHeader(500) 
	   json.NewEncoder(res).Encode(errParseURl.Error())
	   return
   }
   
   errDB := GetAllNewsDeleted(N.DB,&_new ,&_news)		
   
   if errDB != nil {
	   res.Header().Set("Content-Type", "application/json")
	   res.WriteHeader(500)
	   json.NewEncoder(res).Encode(errDB.Error())
	   return
   }
   
   res.Header().Set("Content-Type", "application/json")
   res.WriteHeader(200)
   json.NewEncoder(res).Encode(_news)
   return
}

func (N *NewsService)Create(res http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)
	var _news NewsData
	//var _Topics []Migrate.Topics
	err := decoder.Decode(&_news) 
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	defer req.Body.Close()
	err = InsertNews(N.DB, &_news)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500) 
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(201) 
	json.NewEncoder(res).Encode(_news)
} 

func (N  *NewsService) Uploadimage(res http.ResponseWriter, req *http.Request){
	req.ParseMultipartForm(32 << 20)
	file, handler, err := req.FormFile("banner")
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(500) 
		json.NewEncoder(res).Encode(err)
		return

	}
	defer file.Close()
	var filename = time.Now().UTC().String()+"-"+handler.Filename
	f, err := os.OpenFile("./public/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200) 
	json.NewEncoder(res).Encode(filename)
}

func (N *NewsService)Update(res http.ResponseWriter, req *http.Request){
	
	var _news NewsData
	var id = req.URL.Query().Get("id")
	_id, err := strconv.Atoi(id)
	err= GetNews(N.DB, _id , &_news)
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&_news) 
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	defer req.Body.Close()
	err = UpdateNews(N.DB,_id, &_news)
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	 }
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200) 
	json.NewEncoder(res).Encode(_news)
}

func (N *NewsService)Delete(res http.ResponseWriter, req *http.Request){
	var _news NewsData
	var id = req.URL.Query().Get("id")
	var message Config.Response
	_id, err := strconv.Atoi(id)
	err = GetNews(N.DB, _id , &_news)
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	}
	err = DeletedNews(N.DB , &_news)
	if err != nil {
		res.WriteHeader(500)
		json.NewEncoder(res).Encode(err.Error())
		return
	 }
	 res.Header().Set("Content-Type", "application/json")
	 res.WriteHeader(200) 
	 message.Message = "news delete"
	 json.NewEncoder(res).Encode(message)
}