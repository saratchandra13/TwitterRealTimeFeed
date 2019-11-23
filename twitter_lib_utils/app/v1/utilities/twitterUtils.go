package utilities

import(
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"twitter_stream/twitter_lib_utils/app/v1/models"
)

const(
	ConsumerKey = "hKLesXRM3djYfkrrNwTN2lWgr"
	ConsumerSecret = "xtM9CQVXNxfpxYuxTNFN7LWYWHOnmqE1ykAES8lPqCIpp1IEYV"
	AccessToken = "869240018664964097-yBzhazJ7S8UDngf6uufa0iejwtwTGnu"
	AccessSecret = "RAQpQIdNrHaT0RTosDPjunTIBudcfOYzzQIYrftVYOyBA"
)

var TwitterClient *twitter.Client

type TwitterSearchUtils struct {

}

type TweetSearcher interface {
	GetTweetsFromTwitter(context *gin.Context,tweetsearcher TweetSearcher,AccountName string,TweetChannel chan models.TweetModel)
	ListenForOsSignalsForExit(chan os.Signal)
}

func init(){

	fmt.Println("Initializing twitter Client")

	config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	token := oauth1.NewToken(AccessToken, AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	TwitterClient = twitter.NewClient(httpClient)

	fmt.Println("Initializing successful")

	return
}

func (twitterSearchUtilsVariable TwitterSearchUtils) GetTweetsFromTwitter(context *gin.Context,TweetSearcher TweetSearcher,AccountName string,TweetChannel chan models.TweetModel){

	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic Caught! Error in GetTweetsFromTwitter : " + "General Error: " + fmt.Sprint(r) + " and stacktrace is" + string(debug.Stack())
			log.Print(ErrorString)
			err := errors.New(ErrorString)
			TweetChannel <- models.TweetModel{TweetText:"",ErrorString:err}
			close(TweetChannel)
		}
	}()

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		TweetChannel <- models.TweetModel{TweetText:tweet.Text,Sender:tweet.User.Name,RetweetedCount:tweet.RetweetCount,ReplyCount:tweet.ReplyCount,CreatedAt:tweet.CreatedAt}
	}

	// FILTER
	filterParams := &twitter.StreamUserParams{
		Track:         []string{AccountName},
		StallWarnings: twitter.Bool(true),
	}

	if UserStream, err := TwitterClient.Streams.User(filterParams);err!=nil{
		TweetChannel <- models.TweetModel{ErrorString:errors.New("fatal error while creating a stream")}
		close(TweetChannel)
		return
	}else{
		go demux.HandleChan(UserStream.Messages)

		ch := make(chan os.Signal)
		TweetSearcher.ListenForOsSignalsForExit(ch)

		if _,ok := <-ch;ok{
			UserStream.Stop()
			close(TweetChannel)
		}
	}
}

func (twitterSearchUtilsVariable TwitterSearchUtils)  ListenForOsSignalsForExit(ch chan os.Signal) {
	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
}