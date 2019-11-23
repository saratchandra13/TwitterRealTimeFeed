// @ApplicationTitle Twitter Stream Search
// @APIDescription Try out Different Twitter Search APIs
// @Contact saratchandra13@gmail.com

package main

import(
	"github.com/gin-gonic/gin"
	"log"
	SearchRouter "twitter_stream/twitter_search/app/v1/router"

)

func main(){
	mainRouter := gin.Default()

	// add all the routes present in the search here.
	TwitterSearchRouterGroup := mainRouter.Group("/search")
	SearchRouter.AddTwitterStreamApisRoutes(TwitterSearchRouterGroup)

	// can add more functionalities here.

	if err := mainRouter.Run(":8080");err!=nil{
		log.Fatal("Fatal Error! Error while starting the server. Shutting Down the server. Goodbye From Server. :( ")
	}else{
		log.Println("Hi! From the server, :D")
	}

}


