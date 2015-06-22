package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


type Hashtag struct {
	Text    string `json:"text"`
	Indices [2]int `json:"indices"`
}

type Size struct {
	Width  int    `json:"w"`
	Height int    `json:"h"`
	Resize string `json:"resize"`
}

type Sizes struct {
	Thumb  Size `json:"thumb"`
	Large  Size `json:"large"`
	Medium Size `json:"medium"`
	Small  Size `json:"small"`
}

type Media struct {
	Id             int64  `json:"id"`
	Indices        [2]int `json:"indices"`
	Url            string `json:"url"`
	DisplayUrl     string `json:"display_url"`
	ExpandedUrl    string `json:"expanded_url"`
	MediaUrlHttps  string `json:"media_url_https"`
	Sizes          Sizes
	SourceStatusId int64  `json:"source_status_id"`
	Type           string `json:"type"`
}

type Url struct {
	Url         string `json:"url"`
	DisplayUrl  string `json:"display_url"`
	ExpandedUrl string `json:"expanded_url"`
	Indices     [2]int `json:"indices"`
}

type UserMention struct {
	Id         int64  `json:"id"`
	Indices    [2]int `json:"indices"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type Entities struct {
	Hashtags     []Hashtag     `json:"hashtags"`
	Media        []Media       `json:"media"`
	Urls         []Url         `json:"urls"`
	UserMentions []UserMention `json:"user_mentions"`
}

type User struct {
	Id                   int64  `json:"id"`
	Name                 string `json:"name"`
	ScreenName           string `json:"screen_name"`
	CreatedAt            string `json:"created_at"`
	Description          string `json:"description"`
	FavouritesCount      int    `json:"favourites_count"`
	FollowersCount       int    `json:"followers_count"`
	FriendsCount         int    `json:"friends_count"`
	GeoEnabled           bool   `json:"geo_enabled"`
	Lang                 string `json:"lang"`
	ListedCount          int    `json:"listed_count"`
	Location             string `json:"location"`
	ProfileImageUrlHttps string `json:"profile_image_url_https"`
	Status               *Tweet `json:"status"`
	StatusesCount        int    `json:"statuses_count"`
}

type Coordinates struct {
	Coordinates []float64 `json:"coordinates"`
	Type string `json:"type"`
}

type Tweet struct {
	Id                  int64        `json:"id"`
	Text                string       `json:"text"`
	Entities            Entities     `json:"entities"`
	RetweetCount        int          `json:"retweet_count"`
	Source              string       `json:"source"`
	Coordinates         *Coordinates `json:"coordinates"`
	User                User         `json:"user"`
	CreatedAt           string       `json:"created_at"`
	FavoriteCount       int          `json:"favorite_count"`
	InReplyToScreenName string       `json:"in_reply_to_screen_name"`
	InReplyToStatusId   int64        `json:"in_reply_to_status_id"`
	InReplyToUserId     int64        `json:"in_reply_to_user_id"`
	Lang                string       `json:"lang"`
	RetweetedStatus     *Tweet       `json:"retweeted_status"`
}


func main() {

	// Create new tweet
	gen_tweet := &Tweet{
		Text: "hello, go",
		User: User{
			ScreenName: "User",
		},
		Entities: Entities{
			Hashtags: []Hashtag{
				Hashtag{
					Text:    "#go",
					Indices: [2]int{1, 2},
				},
				Hashtag{
					Text:    "#golang",
					Indices: [2]int{4, 7},
				},
			},
		},
	}
	
	// Encode and print it
	gen_tweet_encoded, _ := json.Marshal(gen_tweet)
	fmt.Println(string(gen_tweet_encoded))

	// Read file
	data, e := ioutil.ReadFile("tweet.json")
	if e != nil {
		panic(e)
	}

	// Parse tweet into our structure
	tweet := &Tweet{}
	json.Unmarshal(data, &tweet)

	// Show simple fields
	fmt.Println(tweet.CreatedAt)
	fmt.Println(tweet.User.ScreenName)
	fmt.Println(tweet.RetweetCount)
	fmt.Println(tweet.Text)
	
	// Print first media if exists
	if len(tweet.Entities.Media) > 0{
		fmt.Println(tweet.Entities.Media[0].Type)	
	}
	
	// If retweet loaded, show source text
	rt := tweet.RetweetedStatus
	if rt != nil {
		fmt.Println(rt.Text)
	}

}
