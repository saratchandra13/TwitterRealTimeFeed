package models

type SearchResponse struct {
	SearchWord string `json:"searchWord"`
	Tweet string `json:"tweet"`
	RetweetCount int `json:"rt_count"`
	ReplyCount int `json:"reply_count"`
	Sender string `json:"senderName"`
	TweetTime string `json:"tweetTime"`
	Error error `json:"error"`
}
