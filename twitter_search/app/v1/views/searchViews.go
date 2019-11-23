package views


import(
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"errors"
	"twitter_stream/twitter_lib_utils/app/v1/utilities"
	"twitter_stream/twitter_search/app/v1/models"
	TwitterModel "twitter_stream/twitter_lib_utils/app/v1/models"
)


type RealTimeTweets struct {

}

type RealTimeTweetsListener interface {
	GetRealTimeTweets(context *gin.Context,RealTimeTweetsListener RealTimeTweetsListener,SearchWord string,ResponseChannel chan models.SearchResponse)
}

func (SearchViewsVariable RealTimeTweets) GetRealTimeTweets(context *gin.Context,RealTimeTweetsListener RealTimeTweetsListener,SearchWord string,ResponseChannel chan models.SearchResponse){

	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic Caught! Error in GetRealTimeTweets : " + "General Error: " + fmt.Sprint(r) + " and stacktrace is" + string(debug.Stack())
			log.Print(ErrorString)
			err := errors.New(ErrorString)
			ResponseChannel <- models.SearchResponse{SearchWord:SearchWord,Error:err}
			close(ResponseChannel)
		}
	}()

	var TwitterUtils utilities.TwitterSearchUtils
	TweetChannel := make(chan TwitterModel.TweetModel)

	go TwitterUtils.GetTweetsFromTwitter(context,TwitterUtils,SearchWord,TweetChannel)

	for {
		if Tweet, ok := <-TweetChannel; ok {
			ResponseChannel <- models.SearchResponse{Sender: Tweet.Sender, SearchWord: SearchWord, Tweet: Tweet.TweetText, TweetTime: Tweet.CreatedAt}
		} else {
			fmt.Println("Closing the realtime tweets, Thankyou")
			close(ResponseChannel)
			break
		}
	}
	return
}