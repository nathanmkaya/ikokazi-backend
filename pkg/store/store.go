package store

import "github.com/nathanmkaya/ikokazi-backend/pkg/model"

type Store interface {
	AddTweet(tweet model.Tweet)
	GetTweet(id int)
	GetAllTweets()
	GetTweetsByKeyword(...string)
	AddDoc(interface{}) error
	GetDoc()
	OnNewTweet() chan interface{}
}
