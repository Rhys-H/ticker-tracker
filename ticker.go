package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/joho/godotenv"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	os.Setenv(common.EnvApiKeyID, os.Getenv("API_KEY"))
	os.Setenv(common.EnvApiSecretKey, os.Getenv("API_SECRET"))

	alpaca.SetBaseUrl("https://paper-api.alpaca.markets")
}

func main() {
	// Alpaca API currently not used
	alpacaClient := alpaca.NewClient(common.Credentials())
	acct, err := alpacaClient.GetAccount()
	if err != nil {
		panic(err)
	}
	fmt.Println(*acct)

	posts := dailyTopPosts("wallstreetbets")
	mentions := checkMentions(posts, "TSLA")

	if mentions == nil {
		fmt.Println("No mentions of supplied stocks")
	}
}

func dailyTopPosts(sub string) []*reddit.Post {
	client, _ := reddit.NewReadonlyClient()
	posts, _, err := client.Subreddit.TopPosts(context.Background(), sub, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		Time: "today",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received %d posts.\n", len(posts))

	return posts
}

func checkMentions(posts []*reddit.Post, stock string) []string {
	var mentions []string

	for _, post := range posts {
		if strings.Contains(post.Title, stock) {
			mentions = append(mentions, post.Title)
			fmt.Println(post.Title)
		}
	}

	return mentions
}
