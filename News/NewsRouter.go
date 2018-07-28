package News
import (
	"github.com/gorilla/mux"
	"bytes"
	"github.com/jinzhu/gorm"
)

type  NewsRouter struct{
	DB *gorm.DB
}



func (N * NewsRouter) AddSignHandler(r *mux.Router, db *gorm.DB){

	var nService NewsService
	nService.setGorm(db)
	// swagger:route GET /news News ListNews
	//
	// Lists all news.
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

	// swagger:parameters News ListNews
	type newsparams struct {
		// in: query
		ID int `json:"id"`
		Title string `json:"title"`
		Is_draf bool `json:"is_draf"`
		Is_publish bool `json:"is_publish"`
	}
	r.HandleFunc("/news",nService.Get).Methods("Get")
	
	// swagger:route GET /news/deleted News ListDeletedNews
	//
	// Lists all Deleted news.
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
    // swagger:parameters News ListDeletedNews
	type newsparamsdeleted struct {
		// in: query
		ID int `json:"id"`
		Title string `json:"title"`
		Is_draf bool `json:"is_draf"`
		Is_publish bool `json:"is_publish"`
	}	 
	r.HandleFunc("/news/deleted",nService.GetDeletedNews).Methods("GET")
	
	// swagger:route GET /news/topic News ListtopicNews
	//
	// Lists all  news. by topics
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
    // swagger:parameters News ListtopicNews
	type newsparamstopics struct {
		// id topics
		// in: query
		ID int `json:"id"`
	}
	r.HandleFunc("/news/topic", nService.GetbyTopics).Methods("GET")
	
	// swagger:route POST /news News createnews
	//
	// create news.
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
	// swagger:parameters News createnews
	type newsparamsnew struct {
		//   description: 
	    //   json format to send
		//   {
		//     "title": "title",
		//     "body": "lorem ipsum",
		//     "banner": "banner.png",
		//     "datePublish": "2018-07-27T17:58:39Z",
		//     "is_draf": true,
		//     "is_publish": false,
		//     "Topics": [{
        //	"topic_name":"coba"
		//	}]
		// }
		//topics it's optional for create with new topics
		// in: body
		 Body string   
	}
	r.HandleFunc("/news",nService.Create).Methods("POST")
	// swagger:route GET /images/{filename} News getbanner
	//
	// getbanner.
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
	// swagger:parameters News getbanner
	type getbanner struct {
		
		// In: path
		Filename string `json:"filename"`   
	}

	
	// swagger:route POST /news/uploadbanner News uploadbanner
	//
	// uploadbanner.
	//
	// This will show all recorded people.
	//
	//     Consumes:
	//     - multipart/form-data
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       201:
	// swagger:parameters News uploadbanner
	type newsparamsnewbanner struct {
		

		// in: formData
		// swagger:file
		Banner *bytes.Buffer `json:"banner"`   
	}
	
	r.HandleFunc("/news/uploadbanner",nService.Uploadimage).Methods("POST")
	
	// swagger:route PUT/news News updatenews
	//
	// update news.
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
	// swagger:parameters News updatenews
	type newsparamsupdate struct {
		//in:query
		ID int `json:"id"`

		//   description: 
	    //   json format to send
		//   {
		//     "title": "title",
		//     "body": "lorem ipsum",
		//     "banner": "banner.png",
		//     "datePublish": "2018-07-27T17:58:39Z",
		//     "is_draf": true,
		//     "is_publish": false,
		//     "Topics": [{
        //	"topic_name":"coba"
		//	}]
		// }
		//topics it's optional for create with new topics
		//note add topics thats mean to create new topics
		// in: body
		 Body string   
	}
	r.HandleFunc("/news",nService.Update).Methods("PUT")

	// swagger:route DELETE /news News delete
	//
	// delete news.
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
	// swagger:parameters News delete
	type newsdelete struct {
		// in: query
		ID int `json:"id"`
	}
	r.HandleFunc("/news",nService.Delete).Methods("Delete")
}
