package scraper

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

// Global context and cancel function to keep browser open
var ctx context.Context
var cancel context.CancelFunc

// InitializeHeadlessChrome starts a headless Chrome browser and sets up the context
func InitializeHeadlessChrome() {
	// Create a new context and cancel function
	ctx, cancel = chromedp.NewContext(context.Background())
	// Set timeout for browser operations
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
}

// CloseHeadlessChrome closes the browser when done
func CloseHeadlessChrome() {
	if cancel != nil {
		cancel() // Call the cancel function to close the browser session
	}
}

// GetVerifiedFollowersLink extracts the verified followers href from a Twitter profile
func GetVerifiedFollowersLink(profileURL string) (string, error) {
	var href string

	err := chromedp.Run(ctx,
		chromedp.Navigate(profileURL),
		chromedp.Sleep(5*time.Second), // wait for dynamic content to load
		chromedp.AttributeValue(`//a[contains(@href, "/verified_followers")]`, "href", &href, nil),
	)

	if err != nil {
		return "", err
	}

	if href == "" {
		return "", fmt.Errorf("verified followers link not found on %s", profileURL)
	}

	return href, nil
}

// FetchAnotherInfoForProfile extracts additional information (e.g., number of tweets) from the same profile
func FetchAnotherInfoForProfile(profileURL string) (string, error) {
	var tweetCount string

	err := chromedp.Run(ctx,
		chromedp.Navigate(profileURL),
		chromedp.Sleep(5*time.Second), // wait for dynamic content to load
		chromedp.Text(`//div[contains(@class, "tweet-count")]`, &tweetCount, chromedp.NodeVisible),
	)

	if err != nil {
		return "", err
	}

	return tweetCount, nil
}
