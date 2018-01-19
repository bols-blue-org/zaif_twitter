package main

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

const (
	consumerKey       = ""
	consumerSecret    = ""
	accessToken       = ""
	accessTokenSecret = ""
)

func init() {

}
func Tweet(message string) {
	// ConsumerKeyのセット
	anaconda.SetConsumerKey(consumerKey)

	// ConsumerSecretのセット
	anaconda.SetConsumerSecret(consumerSecret)

	// AccessTokenとAccessTokenSecretのセット
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	v := url.Values{}

	// ツイートする
	_, err := api.PostTweet(message, v)

	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	Tweet("昨日の頑張りの結果はこちらです！\n素晴らしいですか?")
}
