package main

import (
	"flag"
	"fmt"
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
	} else {
		// Perform a basic scan if the user does not provide any instructions.
		basicScan = true
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
	} else if !(basicWebsiteUsed == resp.Request.URL.String()) {
		log.Println("Error: ", basicWebsiteUsed)
	} else {
		fmt.Println("Passed: ", basicWebsiteUsed)
	}
	_ = resp
}

func advancedNetworkCheck() {
	// Lists of services to test
	websiteTestList := []string{
		"https://www.google.com",
		"https://github.com",
		"https://www.af.mil",
		"https://www.facebook.com",
		"https://www.cia.gov",
		"https://www.amazon.com",
		"https://americorps.gov",
		"https://www.youtube.com",
		"https://www.yahoo.com",
		"https://www.apple.com",
		"https://www.nsa.gov",
		"https://zoom.us",
		"https://www.reddit.com",
		"https://www.redtube.com",
		"https://www.xtube.com",
		"https://www.porn.com",
		"https://www.fbi.gov",
		"https://www.omct.org",
		"https://www.msn.com",
		"https://www.netflix.com",
		"https://www.bing.com",
		"https://www.microsoft.com",
		"https://www.cloudflare.com",
		"https://www.ebay.com",
		"https://www.instagram.com",
		"https://chaturbate.com",
		"https://wfrtds.org",
		"https://www.who.int",
		"https://www.tunnelbear.com",
		"https://www.spacex.com",
		"https://www.xswiper.com",
		"https://www.pornhub.com",
		"https://www.youporn.com",
		"https://4genderjustice.org",
		"https://www.xvideos.com",
		"https://www.tesla.com",
		"https://www.privateinternetaccess.com",
		"https://www.prolife.com",
		"https://www.office.com",
		"https://tinder.com",
		"https://www.hrw.org",
		"https://www.twitch.tv",
		"https://www.bbc.com",
		"https://www.wikipedia.org",
		"https://www.usa.gov",
		"https://bumble.com",
		"https://www.academyadmissions.com",
		"https://www.state.gov",
		"https://www.tsa.gov",
		"https://www.whitehouse.gov",
		"https://www.usds.gov",
		"https://aal.army",
		"https://www.federalreserveeducation.org",
		"https://armyfuturescommand.com",
		"https://www.nasa.gov",
	}
	// Start the test
	for i := 0; i < len(websiteTestList); i++ {
		resp, err := http.Get(websiteTestList[i])
		if err != nil {
			log.Println("Failed: ", websiteTestList[i])
		} else if !(websiteTestList[i] == resp.Request.URL.String()) {
			log.Println("Error: ", websiteTestList[i])
		}
		//else {	//fmt.Println("Passed: ", websiteTestList[i])	}
		_ = resp
	}
}
