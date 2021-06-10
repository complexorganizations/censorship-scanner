package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
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
		log.Fatal("Error: It's impossible to avoid doing a basic and advanced scan.")
	}
	// Only perform one scan at a time.
	if basicScan && advancedScan {
		log.Fatal("Error: A basic and advanced scan cannot be performed at the same time.")
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
	// Send the http request and check for any certificates.
	sendTheRequest(basicWebsiteUsed)
	fmt.Println("Private IP:", getCurrentPrivateIP())
	fmt.Println("Public IP:", getCurrentPublicIP())
}

func advancedNetworkCheck() {
	// Lists of services to test
	websiteTestList := []string{
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.pinkcupid.com",
		"https://www.baidu.com",
		"https://www.yahoo.com",
		"https://www.gumlet.com",
		"https://www.keycdn.com",
		"https://www.cachefly.com",
		"https://www.netlify.com",
		"https://www.psiphon3.com",
		"https://riseup.net",
		"https://helm.sh",
		"https://goharbor.io",
		"https://www.fastly.com",
		"https://www.digitalocean.com",
		"https://www.vultr.com",
		"https://www.synology.com",
		"https://www.jaegertracing.io",
		"https://ovhcloud.com",
		"https://www.kraken.com",
		"https://www.akamai.com",
		"https://www.rackspace.com",
		"https://www.openstack.org",
		"https://tv.youtube.com",
		"https://www.paramountplus.com",
		"https://vuejs.org",
		"https://www.kernel.org",
		"https://996.icu",
		"https://www.tensorflow.org",
		"https://fontawesome.com",
		"https://getbootstrap.com",
		"https://redis.io",
		"https://www.amazon.com",
		"https://www.hostwinds.com",
		"https://www.netrepid.com",
		"https://cloud.ionos.com",
		"https://www.wikipedia.org",
		"https://www.freecodecamp.org",
		"https://tubitv.com",
		"https://www.disneyplus.com",
		"https://www.sling.com",
		"https://rsf.org",
		"https://www.linode.com",
		"https://www.fubo.tv",
		"https://discord.com",
		"https://www.philo.com",
		"https://www.hbomax.com",
		"https://www.plex.tv",
		"https://www.kanopy.com",
		"https://www.pbs.org",
		"https://weather.com",
		"https://www.qq.com",
		"https://www.google.co.in",
		"https://twitter.com",
		"https://www.peacocktv.com",
		"https://outlook.live.com",
		"https://www.osce.org",
		"https://sourceforge.net",
		"https://ultrasurf.us",
		"https://nypost.com",
		"https://martus.org",
		"https://www.ecns.cn",
		"https://www.torproject.org",
		"https://www.taobao.com",
		"https://azure.microsoft.com",
		"https://www.bing.com",
		"https://www.instagram.com",
		"https://www.weibo.com",
		"https://signal.org",
		"https://www.messenger.com",
		"https://telegram.org",
		"https://www.sina.com.cn",
		"https://www.linkedin.com",
		"https://www.yahoo.co.jp",
		"https://www.msn.com",
		"https://vk.com",
		"https://operatorframework.io",
		"https://www.google.de",
		"https://yandex.ru",
		"https://www.hao123.com",
		"https://www.google.co.uk",
		"https://spiffe.io",
		"https://www.wsj.com",
		"https://www.bloomberg.com",
		"https://cloudevents.io",
		"https://news.google.com",
		"https://www.theguardian.com",
		"https://prometheus.io",
		"https://etcd.io",
		"https://www.openpolicyagent.org",
		"https://cortexmetrics.io",
		"https://www.cncf.io",
		"https://opentracing.io",
		"https://kubernetes.io",
		"https://coredns.io",
		"https://cloud.google.com",
		"https://grpc.io",
		"https://thanos.io",
		"https://www.ea.com",
		"https://aws.amazon.com",
		"https://www.hashicorp.com",
		"https://www.envoyproxy.io",
		"https://www.sap.com",
		"https://www.businessinsider.com",
		"https://www.samsung.com",
		"https://www.cisco.com",
		"https://www.netflix.com",
		"https://www.hulu.com",
		"https://www.adobe.com",
		"https://www.workday.com",
		"https://mega.io",
		"https://www.servicenow.com",
		"https://www.itemfix.com",
		"https://duckduckgo.com",
		"https://www.prajwalkoirala.com",
		"https://time.com",
		"https://flickr.com",
		"https://www.vmware.com",
		"https://www.plurk.com",
		"https://www.nutanix.com",
		"https://slack.com",
		"https://soundcloud.com",
		"https://gitlab.com",
		"https://www.fox.com",
		"https://www.nbc.com",
		"https://www.ibm.com",
		"https://www.oracle.com",
		"https://www.salesforce.com",
		"https://www.reddit.com",
		"https://www.ebay.com",
		"https://www.google.fr",
		"https://www.t.co",
		"https://www.tmall.com",
		"https://www.google.com.br",
		"https://www.360.cn",
		"https://www.sohu.com",
		"https://www.sogou.com",
		"https://ok.ru",
		"https://www.amazon.co.jp",
		"https://www.pinterest.com",
		"https://www.google.it",
		"https://www.antgroup.com",
		"https://mail.ru",
		"https://amp.dev",
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
		"https://www.foxnews.com",
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
		"https://www.quora.com",
		"https://www.fandom.com",
		"https://www.craigslist.org",
		"https://www.amazon.de",
		"https://www.nicovideo.jp",
		"https://www.google.pl",
		"https://www.naver.com",
		"https://www.soso.com",
		"https://www.bilibili.com",
		"https://www.dropbox.com",
		"https://bitly.com",
		"https://www.cricbuzz.com",
		"https://www.outbrain.com",
		"https://www.pixnet.net",
		"https://www.taboola.com",
		"https://www.alibaba.com",
		"https://golang.org",
		"https://www.aliexpress.com",
		"https://www.booking.com",
		"https://www.googleusercontent.com",
		"https://www.google.com.au",
		"https://www.popads.net",
		"https://www.cntv.cn",
		"https://www.tribunnews.com",
		"https://www.zhihu.com",
		"https://www.amazon.co.uk",
		"https://www.diply.com",
		"https://www.coccoc.com",
		"https://www.eff.org",
		"https://www.pornmd.com",
		"https://www.yahoo.com",
		"https://www.cnn.com",
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
	// Send the http request and see if certificates are valid.
	for i := 0; i < len(uniqueDomains); i++ {
		sendTheRequest(uniqueDomains[i])
	}
	fmt.Println("Private IP:", getCurrentPrivateIP())
	fmt.Println("Public IP:", getCurrentPublicIP())
}

// Make all the array unique
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

// Obtain the private IP address of the system.
func getCurrentPrivateIP() []net.IP {
	hostName, err := os.Hostname()
	handleErrors(err)
	getIP, err := net.LookupIP(hostName)
	handleErrors(err)
	return getIP
}

// Obtain the public IP address of the system.
func getCurrentPublicIP() []string {
	var foundIP []string
	url := "https://checkip.amazonaws.com"
	// Verify that the urls are correct.
	if validURL(url) {
		// All insecure http requests are blocked.
		if !strings.Contains(url, "http://") {
			response, err := http.Get(url)
			handleErrors(err)
			body, err := io.ReadAll(response.Body)
			handleErrors(err)
			defer response.Body.Close()
			regex := regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)
			foundIP = regex.FindAllString(string(body), -1)
		}
	}
	return foundIP
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

// send all the request
func sendTheRequest(url string) {
	// Verify that the urls are correct.
	if validURL(url) {
		// All insecure http requests are blocked.
		if !strings.Contains(url, "http://") {
			resp, err := http.Get(url)
			if err != nil {
				log.Println("Censored URL: ", url)
			} else if !(url == resp.Request.URL.String()) {
				//log.Println("Error URL: ", url)
				validateSSLCert(url)
			} else {
				//fmt.Println("Valid URL: ", url)
				validateSSLCert(url)
			}
		}
	}
}

// Validate all the SSL Certs
func validateSSLCert(hostname string) {
	// Take a look at the URL and parse it.
	parsedURL, err := url.Parse(hostname)
	handleErrors(err)
	// obtain the domain name
	parsedHostname := fmt.Sprint(parsedURL.Hostname())
	// verify the ssl
	callTCP, err := tls.Dial("tcp", parsedHostname+":443", nil)
	if err != nil {
		log.Println("Censored SSL:", parsedHostname)
	}
	callTCP.Close()
	err = callTCP.VerifyHostname(parsedHostname)
	if err != nil {
		//log.Println("Error SSL:", parsedHostname)
	} else {
		//fmt.Println("Valid SSL: ", parsedHostname)
	}
}
