package index

import (
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/nathanmkaya/ikokazi-backend/pkg/model"
	"log"
)

type Algolia struct {
	Config *model.Config
	*search.Client
	Name string
}

func (a *Algolia) IndexTweet(tweet interface{}) {
	_, err := a.Client.InitIndex(a.Name).SaveObject(tweet)
	if err != nil {
		log.Fatalln(err)
	}
}

func NewAlgoliaIndex(config *model.Config, name string) Index {
	client := search.NewClient(config.AlgoliaAppID, config.AlgoliaAPIKey)
	return &Algolia{
		Config: config,
		Client: client,
		Name:   name,
	}
}
