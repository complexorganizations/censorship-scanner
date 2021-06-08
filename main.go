package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) > 1 {
	} else {
		// fmt.Println("A comprehensive scan of your network takes less than 5 seconds with network tester.")
	}
}

func main() {
	basicNetworkCheck()
}

func basicNetworkCheck() {
	basicWebsiteUsed := "https://www.example.com"
	resp, err := http.Get(basicWebsiteUsed)
	if err != nil {
		log.Print("Failure: ", basicWebsiteUsed)
	}
	_ = resp
	advancedNetworkCheck()
}

func advancedNetworkCheck() {
	// Lists of services to test
	websiteTestList := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
		"https://www.yahoo.com",
		"https://www.apple.com",
		"https://zoom.us",
		"https://www.reddit.com",
		"https://www.redtube.com",
		"http://www.xtube.com",
		"http://www.porn.com",
		"https://www.msn.com",
		"https://www.netflix.com",
		"https://www.bing.com",
		"https://www.microsoft.com",
		"https://www.cloudflare.com",
		"https://www.ebay.com",
		"https://www.instagram.com",
		"https://chaturbate.com",
		"https://www.xswiper.com",
		"http://www.pornhub.com",
		"http://www.youporn.com",
		"http://www.xvideos.com",
		"https://www.office.com",
		"https://www.twitch.tv",
		"https://www.wikipedia.org",
	}
	// Start the test
	for i := 0; i < len(websiteTestList); i++ {
		resp, err := http.Get(websiteTestList[i])
		if err != nil {
			log.Print("Failure: ", websiteTestList[i])
		}
		_ = resp
	}
}
