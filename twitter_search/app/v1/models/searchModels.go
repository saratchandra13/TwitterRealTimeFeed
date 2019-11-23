package models

type SearchResponse struct {
	SearchWord string `json:"searchWord"`
	Tweet string `json:"tweet"`
	Sender string `json:"senderName"`
	TweetTime string `json:"tweetTime"`
	Error error `json:"error"`
}
