package router

import(
	"github.com/gin-gonic/gin"
	"twitter_stream/twitter_search/app/v1/controllers"
)

func AddTwitterStreamApisRoutes(group *gin.RouterGroup){
	group.GET("/tweets/",controllers.GetTwitterTweetsStream)
}