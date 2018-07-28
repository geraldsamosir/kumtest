// The purpose of this application for asigment test from kumparan.
//
//     Schemes: http
//     Host: localhost:5000
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Gerald Halomoan Samosir<geralddeveloper95@gmail.com>
//
//     Consumes:
//     - aplication/json
//
//     Produces:
//     - aplication/json
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/geraldsamosir/kumtest/Config/Migrate"
	"github.com/geraldsamosir/kumtest/Config"
	"github.com/geraldsamosir/kumtest/News"
	"github.com/geraldsamosir/kumtest/Topic"
	"github.com/geraldsamosir/kumtest/News_topics"
	"github.com/jinzhu/gorm"
)

type App struct{
	DB *gorm.DB
}

type Welcome struct { 
	Message string
}

type Notfound struct{
	Welcome
}


func (W *Welcome) Index(res http.ResponseWriter, req *http.Request){	
	W.Message = "rest api kumptest 1.0.0"
	res.WriteHeader(200) 
	json.NewEncoder(res).Encode(W)
}

func(N *Notfound) NotfoundResponse(res http.ResponseWriter, req *http.Request){
	N.Message = "notfound"
	res.WriteHeader(404) 
	json.NewEncoder(res).Encode(N)
}

func (a *App)Initialize(config *Config.Conf) {
	fmt.Println(config.DB.Username)
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = Migrate.DBMigrate(db)
}

func(a *App)Swagger(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
    http.ServeFile(res, req, "swagger.json")
}




func main(){
	var app App
    Config := Config.GetConfig()
	var  _w Welcome 
	var  _nf Notfound 

	//intialize
	app.Initialize(Config)

	var NewsService  News.NewsRouter
	var  TopicService Topic.TopicRouter
	var News_topicsService News_topics.News_topicsRouter
	r := mux.NewRouter()
	r.HandleFunc("/",_w.Index )
	r.HandleFunc("/swagger.json", app.Swagger).Methods("GET")
	r.Handle("/images/{filename}", http.StripPrefix("/images", http.FileServer(http.Dir("./public/"))))
	r.Handle("/docs/{filename}", http.StripPrefix("/docs", http.FileServer(http.Dir("./docs/swagger/dist/") )))
	NewsService.AddSignHandler(r,app.DB)
	TopicService.AddSignHandler(r,app.DB)
	News_topicsService.AddSignHandler(r,app.DB)

	r.NotFoundHandler = http.HandlerFunc(_nf.NotfoundResponse)
	fmt.Println("server run in port 5000")
	http.ListenAndServe("0.0.0.0:5000", r)
}