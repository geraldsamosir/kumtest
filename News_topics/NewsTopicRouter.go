package News_topics
import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

)

type News_topicsRouter struct{

}

func (N * News_topicsRouter) AddSignHandler(r *mux.Router , db *gorm.DB){
	var ntService NewsTopicService
	ntService.setGorm(db)
	
	r.HandleFunc("/newstopics",ntService.Get).Methods("GET")
	r.HandleFunc("/newstopics",ntService.Create).Methods("POST")
	r.HandleFunc("/newstopics",ntService.Update).Methods("PUT")
	r.HandleFunc("/topics",ntService.Delete).Methods("Delete")
}