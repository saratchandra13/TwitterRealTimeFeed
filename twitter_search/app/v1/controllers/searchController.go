package controllers

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"encoding/json"
	"strings"
	"twitter_stream/twitter_search/app/v1/models"
	"twitter_stream/twitter_search/app/v1/views"
)

const(
	Hastag = "#"
	AccountTag = "@"
	BadInputError = "Please search for an Account/HashTag"
)

func GetTwitterTweetsStream(context *gin.Context){
	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic Recoved , Error in GetTwitterTweetsStream and Error = " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			log.Fatal(ErrorString)
		}
	}()

	Flusher, FlusherFlag := context.Writer.(http.Flusher)

	SearchWord := context.Query("source")

	if !strings.HasPrefix(SearchWord,Hastag) && !strings.HasPrefix(SearchWord,AccountTag){
		context.JSON(400,BadInputError)
		return
	}else{
		var TwitterViews views.RealTimeTweets
		TweetsChannel := make(chan models.SearchResponse)

		TwitterViews.GetRealTimeTweets(context,TwitterViews,SearchWord,TweetsChannel)

		for {
			if Tweet,ok := <-TweetsChannel;ok{
				MarshalledResponse,_ := json.Marshal(Tweet)
				_,_ = fmt.Fprintln(context.Writer, string(MarshalledResponse))
			}else{
				_ , _ = fmt.Fprintln(context.Writer,string("Sorry End of Stream."))
			}
			if FlusherFlag{
				Flusher.Flush()
			}
		}

	}
	return
}
