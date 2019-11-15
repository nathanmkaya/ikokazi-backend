package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nathanmkaya/ikokazi-backend/pkg/index"
	"github.com/nathanmkaya/ikokazi-backend/pkg/model"
	"github.com/nathanmkaya/ikokazi-backend/pkg/runner"
	"github.com/nathanmkaya/ikokazi-backend/pkg/store"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := model.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Println(err)
	}
	log.Println(cfg)
	var name = "ikokazike"
	twitterIndex := index.NewAlgoliaIndex(&cfg, name)
	firestore := store.NewStore(context.Background(), &cfg, twitterIndex, name)
	twitterRunner := runner.NewRunner(&cfg, firestore)
	twitterRunner.Stream("Post Malone")

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
}
