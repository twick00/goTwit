package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	//var message string
	message := flag.String("tweet", "", "tweet = whatever message you want to send")
	//message = *flag
	flag.Parse()
	fmt.Println(*message)
	api := twitAuth()
	if ok, err := api.VerifyCredentials(); ok {
		if err != nil {
			log.Fatal(err)
		}
		tweet(*message, api)
		if err != nil {
			log.Panic(err)
		}
	}
}
func tweet(message string, api *anaconda.TwitterApi) {
	v := url.Values{}
	fmt.Println(message)
	if message != "" {
		_, err := api.PostTweet(message, v)
		if err != nil {
			log.Panicln(err)
		} else {
			log.Print("Tweeted message")
		}
	}
}
func readSecret() (string, string, string, string) {
	file, err := os.Open("secret.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var creds []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		creds = append(creds, scanner.Text())
	}
	return creds[0], creds[1], creds[2], creds[3]
}
func twitAuth() *anaconda.TwitterApi {
	key, secret, aKey, aSecret := readSecret()
	anaconda.SetConsumerKey(key)
	anaconda.SetConsumerSecret(secret)
	api := anaconda.NewTwitterApi(aKey, aSecret)
	return api
}
