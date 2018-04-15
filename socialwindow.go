package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"flag"
	"os"
	"log"
	"github.com/coreos/pkg/flagutil"
	"fmt"
	"time"
	"encoding/csv"
	"strconv"
)

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")

	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	client.Accounts.VerifyCredentials(verifyParams)

	f, err := os.Open("D:\\tweets.csv")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return;
	}

	cutoff := time.Now().AddDate(0, -6, 0)
	reader := csv.NewReader(f)
	reader.Read()
	for {
		row, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}

		timestamp, _ := time.Parse("2006-01-02 15:04:05 -0700", row[3])
		if cutoff.Before(timestamp) {
			continue;
		}

		id, _ := strconv.ParseInt(row[0], 10, 64)
		tweet, _, err := client.Statuses.Destroy(id, nil)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("Deleted #%d : %s : %s", tweet.ID, tweet.CreatedAt, row[6])
		fmt.Println()
	}
}