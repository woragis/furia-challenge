package main

import (
	"app/utils/scraper"
	"fmt"
	"log"
)

func main() {
	// Initialize the headless browser once
	scraper.InitializeHeadlessChrome()
	defer scraper.CloseHeadlessChrome() // Ensure we close it when the program ends

	// Example usage: Get verified followers link
	profileURL := "https://twitter.com/FURIA"
	link, err := scraper.GetVerifiedFollowersLink(profileURL)
	if err != nil {
		log.Fatal("Error fetching followers link:", err)
	}
	fmt.Println("Verified followers link:", link)

	// Example usage: Fetch another profile info (e.g., tweet count)
	tweetCount, err := scraper.FetchAnotherInfoForProfile(profileURL)
	if err != nil {
		log.Fatal("Error fetching tweet count:", err)
	}
	fmt.Println("Number of tweets:", tweetCount)
}
