package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
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
	// Close the program if there are no scan kinds.
	if !basicScan && !advancedScan {
		log.Fatal("Error: It is not possible to not perform a basic and an advanced scan at the same time.")
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
	if validURL(basicWebsiteUsed) {
		resp, err := http.Get(basicWebsiteUsed)
		if err != nil {
			log.Println("Failed: ", basicWebsiteUsed)
		} else if !(basicWebsiteUsed == resp.Request.URL.String()) {
			log.Println("Error: ", basicWebsiteUsed)
		} else {
			fmt.Println("Passed: ", basicWebsiteUsed)
		}
	}
	fmt.Println(getCurrentIP())
}

func advancedNetworkCheck() {
	// Lists of services to test
	websiteTestList := []string{
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.baidu.com",
		"https://www.yahoo.com",
		"https://www.amazon.com",
		"https://www.wikipedia.org",
		"https://www.qq.com",
		"https://www.google.co.in",
		"https://twitter.com",
		"https://outlook.live.com",
		"https://www.taobao.com",
		"https://www.bing.com",
		"https://www.instagram.com",
		"https://www.weibo.com",
		"https://www.sina.com.cn",
		"https://www.linkedin.com",
		"https://www.yahoo.co.jp",
		"https://www.msn.com",
		"https://vk.com",
		"https://www.google.de",
		"https://yandex.ru",
		"https://www.hao123.com",
		"https://www.google.co.uk",
		"https://www.reddit.com",
		"https://www.ebay.com",
		"https://www.google.fr",
		"https://www.t.co",
		"https://www.tmall.com",
		"https://www.google.com.br",
		"https://www.360.cn",
		"https://www.sohu.com",
		"https://www.amazon.co.jp",
		"https://www.pinterest.com",
		"https://www.netflix.com",
		"https://www.google.it",
		"https://www.google.ru",
		"https://www.microsoft.com",
		"https://www.google.es",
		"https://www.wordpress.com",
		"https://www.gmw.cn",
		"https://www.tumblr.com",
		"https://www.paypal.com",
		"https://www.blogspot.com",
		"https://imgur.com",
		"https://www.stackoverflow.com",
		"https://www.aliexpress.com",
		"https://hentaihaven.xxx",
		"https://www.naver.com",
		"https://www.ok.ru",
		"https://www.apple.com",
		"https://github.com",
		"https://www.chinadaily.com.cn",
		"https://www.imdb.com",
		"https://www.google.co.kr",
		"https://www.fc2.com",
		"https://www.jd.com",
		"https://www.blogger.com",
		"https://www.163.com",
		"https://www.google.ca",
		"https://www.whatsapp.com",
		"https://www.amazon.in",
		"https://www.office.com",
		"https://www.tianya.cn",
		"https://www.google.co.id",
		"https://www.youku.com",
		"https://www.rakuten.co.jp",
		"https://www.craigslist.org",
		"https://www.amazon.de",
		"https://www.nicovideo.jp",
		"https://www.google.pl",
		"https://www.soso.com",
		"https://www.bilibili.com",
		"https://www.dropbox.com",
		"https://www.outbrain.com",
		"https://www.pixnet.net",
		"https://www.alibaba.com",
		"https://golang.org",
		"https://www.alipay.com",
		"https://www.booking.com",
		"https://www.googleusercontent.com",
		"https://www.google.com.au",
		"https://www.popads.net",
		"https://www.cntv.cn",
		"https://www.zhihu.com",
		"https://www.amazon.co.uk",
		"https://www.diply.com",
		"https://www.coccoc.com",
		"https://www.pornmd.com",
		"https://www.cnn.com",
		"https://www.bbc.co.uk",
		"https://www.twitch.tv",
		"https://www.wikia.com",
		"https://www.google.co.th",
		"https://www.google.com.ph",
		"https://www.doubleclick.net",
		"https://www.onet.pl",
		"https://www.googleadservices.com",
		"https://www.accuweather.com",
		"https://www.googleweblight.com",
		"https://www.answers.yahoo.com",
		"https://www.google.com",
		"https://github.com",
		"https://www.af.mil",
		"https://www.cia.gov",
		"https://americorps.gov",
		"https://www.nsa.gov",
		"https://zoom.us",
		"https://www.redtube.com",
		"https://www.xtube.com",
		"https://www.porn.com",
		"https://www.fbi.gov",
		"https://www.omct.org",
		"https://www.cloudflare.com",
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
		"https://www.rust-lang.org",
		"https://www.privateinternetaccess.com",
		"https://www.prolife.com",
		"https://tinder.com",
		"https://www.hrw.org",
		"https://www.bbc.com",
		"https://nordvpn.com",
		"https://getoutline.org",
		"https://www.usa.gov",
		"https://bumble.com",
		"https://go.dev",
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
	uniqueDomains := makeUnique(websiteTestList)
	// Validate the urls
	for i := 0; i < len(uniqueDomains); i++ {
		if validURL(uniqueDomains[i]) {
			// Start the test
			for i := 0; i < len(uniqueDomains); i++ {
				resp, err := http.Get(uniqueDomains[i])
				if err != nil {
					log.Println("Failed: ", uniqueDomains[i])
				} else if !(websiteTestList[i] == resp.Request.URL.String()) {
					log.Println("Error: ", uniqueDomains[i])
				} else {
					fmt.Println("Passed: ", uniqueDomains[i])
				}
			}
		}
	}
	fmt.Println(getCurrentIP())
}

func makeUnique(randomStrings []string) []string {
	flag := make(map[string]bool)
	var uniqueString []string
	for i := 0; i < len(randomStrings); i++ {
		if !flag[randomStrings[i]] {
			flag[randomStrings[i]] = true
			uniqueString = append(uniqueString, randomStrings[i])
		}
	}
	return uniqueString
}

// Take care of any errors that arise.
func handleErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Obtain the current IP address of the user.
func getCurrentIP() []net.IP {
	hostName, err := os.Hostname()
	handleErrors(err)
	getIP, err := net.LookupIP(hostName)
	handleErrors(err)
	return getIP
}

// Validate the URI
func validURL(uri string) bool {
	validUri, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}
	_ = validUri
	return true
}
