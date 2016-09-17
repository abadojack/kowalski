package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/urfave/cli"
)

func main() {
	var query string
	var count int

	app := cli.NewApp()
	app.Version = "1-alpha"
	app.Name = "kowalski"
	app.Usage = "A twitter client that helps follow users who tweeted a specific phrase"
	app.UsageText = "kowalski -q \"QUERY\" [OPTIONS]"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "query, q",
			Usage:       "Follow users with tweets containing `\"QUERY\"`.",
			Destination: &query,
		},
		cli.IntFlag{
			Name:        "count, c",
			Usage:       "Follow COUNT users.",
			Value:       5,
			Destination: &count,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			app.Name = c.Args()[0]
		}
		if len(query) == 0 {
			errStr := "Usage: " + app.UsageText + "\nTry '" + app.Name + " --help' for more information."
			return errors.New(errStr)
		}

		followUserKeyword(query, count)

		return nil
	}

	app.Run(os.Args)
}

func followUserKeyword(keyword string, count int) {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessKey := os.Getenv("TWITTER_ACCESS_KEY")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	api := getTwitterAPI(consumerKey, consumerSecret, accessKey, accessSecret)

	searchResult, err := api.GetSearch(keyword, nil)
	if err != nil {
		panic(err)
	}

	i := 0
	for {
		for _, tweet := range searchResult.Statuses {
			if i >= count {
				return
			}
			if !tweet.User.Following {
				user, err := api.FollowUser(tweet.User.ScreenName)
				if err != nil {
					log.Panic(err)
				}
				i++

				fmt.Println("[", i, "] ", user.ScreenName, " followed.")
			}
		}
		searchResult, err = searchResult.GetNext(api)
		if err != nil {
			log.Panic(err)
		}
	}
	fmt.Println(i, " users followed.")
}

func getTwitterAPI(consumerKey, consumerSecret, accessKey, accessSecret string) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	return anaconda.NewTwitterApi(accessKey, accessSecret)
}
