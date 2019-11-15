package runner

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/nathanmkaya/ikokazi-backend/pkg/model"
	"github.com/nathanmkaya/ikokazi-backend/pkg/store"
	"log"
)

type TwitterRunner struct {
	Config *model.Config
	Store  store.Store
	*twitter.Client
}

func (t *TwitterRunner) Stream(topic ...string) {
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		t.Store.AddTweet(model.Tweet{
			Tweet:    tweet,
			Keywords: nil,
		})
		log.Printf("Tweet ID:%v Text:%v", tweet.ID, tweet.Text)
	}
	// FILTER
	filterParams := &twitter.StreamFilterParams{
		Track:         topic,
		StallWarnings: twitter.Bool(true),
	}
	stream, err := t.Client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	go demux.HandleChan(stream.Messages)
}

func NewRunner(config *model.Config, store store.Store) Runner {
	clientConfig := oauth1.NewConfig(config.TwitterConsumerKey, config.TwitterConsumerSecret)
	token := oauth1.NewToken(config.TwitterAccessToken, config.TwitterAccessSecret)
	httpClient := clientConfig.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	return &TwitterRunner{
		Config: config,
		Store:  store,
		Client: client,
	}
}
