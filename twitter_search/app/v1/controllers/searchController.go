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

	SearchWord := context.Query("source")
	fmt.Println(SearchWord)

	if !strings.HasPrefix(SearchWord,Hastag) && !strings.HasPrefix(SearchWord,AccountTag){
		fmt.Println("here",SearchWord)
		context.JSON(400,BadInputError)
		return
	}else{
		// http flusher.
		Flusher, FlusherFlag := context.Writer.(http.Flusher)

		var TwitterViews views.RealTimeTweets
		TweetsChannel := make(chan models.SearchResponse,10)

		_,_ = fmt.Fprintln(context.Writer, fmt.Sprint("Starting Stream of tweets for the search word ",SearchWord))
		if FlusherFlag{
			Flusher.Flush()
		}

		go TwitterViews.GetRealTimeTweets(context,TwitterViews,SearchWord,TweetsChannel)

		for {
			if Tweet,ok := <-TweetsChannel;ok{
				MarshalledResponse,_ := json.Marshal(Tweet)
				_,_ = fmt.Fprintln(context.Writer, string(MarshalledResponse))
				if FlusherFlag{
					Flusher.Flush()
				}
			}else{
				_ , _ = fmt.Fprintln(context.Writer,string("Sorry End of Stream! Got a Signal from master to QUIT."))
				if FlusherFlag{
					Flusher.Flush()
					break
				}
			}

		}

	}
	return
}
