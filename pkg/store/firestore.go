package store

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/nathanmkaya/ikokazi-backend/pkg/index"
	"github.com/nathanmkaya/ikokazi-backend/pkg/model"
	"google.golang.org/api/iterator"
	"log"
)

type FireStore struct {
	Config *model.Config
	*firestore.Client
	context.Context
	index      index.Index
	collection string
}

func (store *FireStore) OnNewTweet() chan interface{} {
	out := make(chan interface{})
	snapshotIterator := store.Client.Collection("tweets").Snapshots(store.Context)
	defer snapshotIterator.Stop()
	for {
		snap, err := snapshotIterator.Next()
		if err != nil {
			log.Fatalln(err)
		}
		for _, diff := range snap.Changes {
			log.Println("Data added")
			log.Println(diff.Doc.Data())
			out <- diff.Doc.Data()
			log.Println()
		}
	}
	return out
}

func (store *FireStore) GetDoc() {
	users := store.Collection("users")
	iter := users.Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		fmt.Println(doc.Data())
	}
}

func (store *FireStore) AddDoc(interface{}) error {

	panic("implement me")
}

func (store *FireStore) AddTweet(tweet model.Tweet) {
	_, _ = store.Collection(store.collection).Doc(tweet.IDStr).Set(store.Context, tweet)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	log.Fatalln(err)
	//	panic(err)
	//}
	store.index.IndexTweet(tweet)
}

func (store *FireStore) GetTweet(id int) {
	panic("implement me")
}

func (store *FireStore) GetAllTweets() {
	panic("implement me")
}

func (store *FireStore) GetTweetsByKeyword(...string) {
	panic("implement me")
}

func NewStore(ctx context.Context, config *model.Config, index index.Index, collection string) Store {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	client, _ := app.Firestore(ctx)
	return &FireStore{
		Config:     config,
		Client:     client,
		Context:    ctx,
		index:      index,
		collection: collection,
	}
}
