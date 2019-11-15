package model

import "github.com/dghubble/go-twitter/twitter"

type Tweet struct {
	*twitter.Tweet
	Keywords []string
}
