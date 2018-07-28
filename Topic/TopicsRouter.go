package Topic
import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

)

type  TopicRouter struct{

}

func (N * TopicRouter) AddSignHandler(r *mux.Router , db *gorm.DB){
	var tService TopicService
	tService.setGorm(db)
	// swagger:route GET /topics Topic ListTopics
		//
		// Lists all ListTopics.
		//
		// This will show all recorded ListTopics.
		//
		//     
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//
		//     Schemes: http, https
		//     
		//
		//     Responses:
		//       200: 
		//
		// swagger:parameters Topics ListTopics
		type topicsparams struct {
			// in: query
			ID int `json:"id"`
			Topic_name string `json:"topic_name"`
			Parent_id int `json:"parent_id"`
		}	


	r.HandleFunc("/topics",tService.Get).Methods("GET")
    	// swagger:route POST /topics Topic NewTopic
	//
	// create Topic.
	//
	// This will show all recorded people.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       201:
	// swagger:parameters Topic NewTopic
	type Topicparamsnew struct {
		//   description: 
	    //   json format to send
		// {
		// 	"Topic_name": "asd baru",
		//   "Parent_id":1
		// }
		//parent it's optional for create with new topics
		// in: body
		 Body string   
	}
	r.HandleFunc("/topics",tService.Create).Methods("POST")
	// swagger:route PUT/topics Topic updateTopic
	//
	// update Topic.
	//
	// This will show all recorded people.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       200:
	// swagger:parameters Topic updateTopic
	type newsparamsupdate struct {
		//in:query
		ID int `json:"id"`

		//   description: 
	    //   json format to send
		//   {
		//   "Topic_name":"baru",
		//   "Parent_id":1
		// 	
		// }
		//topics it's optional for create with new topics
		//note add topics thats mean to create new topics
		// in: body
		 Body string   
	}
	r.HandleFunc("/topics",tService.Update).Methods("PUT")

	// swagger:route DELETE /topics Topic deleteTopic
	//
	// delete Topic.
	//
	// This will show all recorded people.
	//
	//     
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//     
	//
	//     Responses:
	//       200: 
	// swagger:parameters News deleteTopic
	type newsdelete struct {
		// in: query
		ID int `json:"id"`
	}
	r.HandleFunc("/topics",tService.Delete).Methods("Delete")
}