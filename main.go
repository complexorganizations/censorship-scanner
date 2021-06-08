package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	basicScan    bool
	advancedScan bool
)

func init() {
	// Decide what type of scan to carry out
	if len(os.Args) > 1 {
		tempBasicScan := flag.Bool("basic", false, "Perform a simple network scan.")
		tempAdvancedScan := flag.Bool("advanced", false, "Perform a thorough network scan.")
		flag.Parse()
		basicScan = *tempBasicScan
		advancedScan = *tempAdvancedScan
	}
	// Only perform one scan at a time.
	if basicScan && advancedScan {
		log.Fatal("Error: It is not possible to perform both a basic and an advanced scan at the same time.")
	}
}

func main() {
	if basicScan {
		basicNetworkCheck()
	} else if advancedScan {
		advancedNetworkCheck()
	}
}

func basicNetworkCheck() {
	basicWebsiteUsed := "https://www.example.com"
	resp, err := http.Get(basicWebsiteUsed)
	if err != nil {
		log.Println("Failed: ", basicWebsiteUsed)
	} else {
		fmt.Println("Passed: ", basicWebsiteUsed)
	}
	_ = resp
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
			log.Println("Failed: ", websiteTestList[i])
		} else {
			fmt.Println("Passed: ", websiteTestList[i])
		}
		_ = resp
	}
}
