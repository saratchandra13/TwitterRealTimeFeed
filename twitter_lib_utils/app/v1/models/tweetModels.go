package models

type TweetModel struct {
	TweetText string
	Sender string
	RetweetedCount int
	ReplyCount int
	CreatedAt string
	ErrorString error
}
