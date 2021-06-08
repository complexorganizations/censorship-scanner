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
	websiteTestList := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.youtube.com",
		"https://www.apple.com",
		"https://www.microsoft.com",
		"https://www.cloudflare.com",
		"https://www.wikipedia.org",
	}
	for i := 0; i < len(websiteTestList); i++ {
		resp, err := http.Get(websiteTestList[i])
		if err != nil {
			log.Print("Failure: ", websiteTestList[i])
		}
		_ = resp
	}
}
