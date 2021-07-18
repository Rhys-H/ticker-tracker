package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	v2 "github.com/alpacahq/alpaca-trade-api-go/v2"
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
	stock := "TSLA"

	posts := dailyTopPosts("wallstreetbets")
	mentions := checkMentions(posts, stock)
	if mentions == nil {
		fmt.Println("No mentions of supplied stocks")
	}

	barset := weeklyStockPrice(stock)
	percentChange := priceMovement(barset, stock)
	fmt.Printf("%s moved %v%% over the last 7 days.\n", stock, percentChange)
}

func dailyTopPosts(sub string) []*reddit.Post {
	client, _ := reddit.NewReadonlyClient()
	posts, _, err := client.Subreddit.TopPosts(context.Background(), sub, &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 100,
		},
		Time: "week",
	})
	if err != nil {
		panic(err)
	}

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

func weeklyStockPrice(stock string) []v2.Bar {
	alpacaClient := alpaca.NewClient(common.Credentials())

	bars := alpacaClient.GetBars(
		stock, v2.Day, v2.Raw, time.Now().Add(-7*24*time.Hour), time.Now().Add(-20*time.Minute), 7)
	var barset []v2.Bar

	for bar := range bars {
		if bar.Error != nil {
			panic(bar.Error)
		}
		barset = append(barset, bar.Bar)
	}

	return barset
}

func priceMovement(barset []v2.Bar, stock string) float64 {
	// See the historical price movement for a given stock
	startPrice := barset[0].Open
	endPrice := barset[len(barset)-1].Close
	percentChange := ((endPrice - startPrice) / startPrice) * 100

	return percentChange
}
