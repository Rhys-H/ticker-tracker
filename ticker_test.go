package main

import (
	"testing"
	"time"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

var testPost1 = &reddit.Post{
	ID:      "i2gvs1",
	FullID:  "t3_i2gvs1",
	Created: &reddit.Timestamp{time.Date(2020, 8, 2, 18, 23, 37, 0, time.UTC)},
	Edited:  &reddit.Timestamp{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},

	Permalink: "/r/test/comments/i2gvs1/this_is_a_title/",
	URL:       "http://example.com",

	Title: "This is a title",

	Likes: reddit.Bool(true),

	Score:            1,
	UpvoteRatio:      1,
	NumberOfComments: 0,

	SubredditName:         "test",
	SubredditNamePrefixed: "r/test",
	SubredditID:           "t5_2qh23",
	SubredditSubscribers:  8278,

	Author:   "v_95",
	AuthorID: "t2_164ab8",
}

var testPost2 = &reddit.Post{
	ID:      "i2gvs1",
	FullID:  "t3_i2gvs1",
	Created: &reddit.Timestamp{time.Date(2020, 8, 2, 18, 23, 37, 0, time.UTC)},
	Edited:  &reddit.Timestamp{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},

	Permalink: "/r/test/comments/i2gvs1/this_is_a_title/",
	URL:       "http://example.com",

	Title: "TEST is the best company ever!!!1!",

	Likes: reddit.Bool(true),

	Score:            1,
	UpvoteRatio:      1,
	NumberOfComments: 0,

	SubredditName:         "test",
	SubredditNamePrefixed: "r/test",
	SubredditID:           "t5_2qh23",
	SubredditSubscribers:  8278,

	Author:   "v_95",
	AuthorID: "t2_164ab8",
}

func TestCheckMentionsReturnsValues(t *testing.T) {
	var testPosts []*reddit.Post
	testPosts = append(testPosts, testPost1, testPost2)

	result := checkMentions(testPosts, "TEST")

	if result == nil {
		t.Fatalf(`checkMentions(testPosts, "TEST") = %q`, result)
	}
}

func TestCheckMentionsReturnsEmpty(t *testing.T) {
	var testPosts []*reddit.Post
	testPosts = append(testPosts, testPost1, testPost2)

	result := checkMentions(testPosts, "NONE")

	if result != nil {
		t.Fatalf(`checkMentions(testPosts, "NONE") = %q`, result)
	}
}
