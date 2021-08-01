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
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	basicScan    bool
	advancedScan bool
	err          error
	wg           sync.WaitGroup
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
		wg.Add(2)
		go advancedNetworkCheck()
		go publicDnsTest()
		torExitNodeTest()
		wg.Wait()
	}
}

func basicNetworkCheck() {
	basicWebsiteUsed := []string{
		"https://www.example.com",
		"https://www.example.net",
		"https://www.example.org",
	}
	sort.Strings(basicWebsiteUsed)
	// Send the http request and check for any certificates.
	for _, basicWebsiteList := range basicWebsiteUsed {
		sendTheRequest(basicWebsiteList)
	}
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
		"https://www.troopmessenger.com",
		"https://www.jaegertracing.io",
		"https://ovhcloud.com",
		"https://www.teamwork.com",
		"https://www.kraken.com",
		"https://www.akamai.com",
		"https://www.atlassian.com",
		"https://www.twilio.com",
		"https://www.mediafire.com",
		"https://www.proworkflow.com",
		"https://www.rackspace.com",
		"https://www.wrike.com",
		"https://bloomfire.com",
		"https://rocket.chat",
		"https://www.openstack.org",
		"https://www.paramountplus.com",
		"https://vuejs.org",
		"https://filestage.io",
		"https://www.kernel.org",
		"https://www.proofhub.com",
		"https://gainapp.com",
		"https://www.skype.com",
		"https://ryver.com",
		"https://www.chanty.com",
		"https://www.tensorflow.org",
		"https://fontawesome.com",
		"https://geti2p.net",
		"https://redis.io",
		"https://monday.com",
		"https://taskworld.com",
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
		"https://www.linkedin.com",
		"https://www.msn.com",
		"https://vk.com",
		"https://operatorframework.io",
		"https://yandex.ru",
		"https://www.hao123.com",
		"https://spiffe.io",
		"https://www.wsj.com",
		"https://www.bloomberg.com",
		"https://cloudevents.io",
		"https://www.theguardian.com",
		"https://aminoapps.com",
		"https://prometheus.io",
		"https://etcd.io",
		"https://thepiratebay.org",
		"https://bandcamp.com",
		"https://www.openpolicyagent.org",
		"https://cortexmetrics.io",
		"https://www.cncf.io",
		"https://www.huffpost.com",
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
		"https://www.tmall.com",
		"https://www.360.cn",
		"https://www.sohu.com",
		"https://www.sogou.com",
		"https://ok.ru",
		"https://www.pinterest.com",
		"https://www.antgroup.com",
		"https://mail.ru",
		"https://amp.dev",
		"https://www.microsoft.com",
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
		"https://www.imdb.com",
		"https://www.fc2.com",
		"https://www.jd.com",
		"https://www.blogger.com",
		"https://www.163.com",
		"https://www.google.ca",
		"https://www.whatsapp.com",
		"https://www.office.com",
		"https://www.tianya.cn",
		"https://www.youku.com",
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
		"https://www.popads.net",
		"https://www.cntv.cn",
		"https://www.tribunnews.com",
		"https://www.zhihu.com",
		"https://www.diply.com",
		"https://www.coccoc.com",
		"https://www.eff.org",
		"https://www.pornmd.com",
		"https://www.yahoo.com",
		"https://www.cnn.com",
		"https://www.twitch.tv",
		"https://www.onet.pl",
		"https://www.accuweather.com",
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
		"https://www.federalreserveeducation.org",
		"https://armyfuturescommand.com",
		"https://www.nasa.gov",
	}
	uniqueDomains := makeUnique(websiteTestList)
	sort.Strings(uniqueDomains)
	// Send the http request and see if certificates are valid.
	for _, uniqueDomainList := range uniqueDomains {
		sendTheRequest(uniqueDomainList)
	}
	fmt.Println("Private IP:", getCurrentPrivateIP())
	fmt.Println("Public IP:", getCurrentPublicIP())
	wg.Done()
}

// To see if you can connect to the Tor network.
func torExitNodeTest() {
	torExitIPs := getTorExitNodes()
	sort.Strings(torExitIPs)
	for i := 0; i < 250; i++ {
		_, err = net.DialTimeout("tcp", torExitIPs[i]+":80", time.Duration(2)*time.Second)
		if err != nil {
			log.Println("Censored TOR:", torExitIPs[i])
		} else {
			fmt.Println("Valid TOR:", torExitIPs[i])
		}
	}
	wg.Done()
}

func publicDnsTest() {
	publicDnsList := []string{
		"8.8.8.8",
		"8.8.4.4",
		"1.1.1.1",
		"1.0.0.1",
		"208.67.222.222",
		"208.67.220.220",
		"9.9.9.9",
		"149.112.112.112",
		"8.26.56.26",
		"8.20.247.20",
		"195.46.39.39",
		"195.46.39.40",
		"77.88.8.8",
		"77.88.8.1",
		"94.140.14.14",
		"94.140.15.15",
		"64.6.64.6",
		"64.6.65.6",
		"80.67.169.40",
		"80.67.169.12",
	}
	sort.Strings(publicDnsList)
	for _, publicDns := range publicDnsList {
		_, err = net.DialTimeout("tcp", publicDns+":53", time.Duration(2)*time.Second)
		if err != nil {
			log.Println("Censored DNS:", publicDns)
		} else {
			fmt.Println("Valid DNS:", publicDns)
		}
	}
	wg.Done()
}

// Make all the array unique
func makeUnique(randomStrings []string) []string {
	flag := make(map[string]bool)
	var uniqueString []string
	for _, content := range randomStrings {
		if !flag[content] {
			flag[content] = true
			uniqueString = append(uniqueString, content)
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
		if !strings.HasPrefix(url, "http://") {
			response, err := http.Get(url)
			handleErrors(err)
			body, err := io.ReadAll(response.Body)
			handleErrors(err)
			response.Body.Close()
			regex := regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)
			foundIP = regex.FindAllString(string(body), -1)
		}
	}
	return foundIP
}

// Obtain the public IP address of all tor exit nodes.
func getTorExitNodes() []string {
	var torExitNodeIPS []string
	url := "https://check.torproject.org/torbulkexitlist"
	// Verify that the urls are correct.
	if validURL(url) {
		// All insecure http requests are blocked.
		if !strings.HasPrefix(url, "http://") {
			response, err := http.Get(url)
			handleErrors(err)
			body, err := io.ReadAll(response.Body)
			handleErrors(err)
			response.Body.Close()
			regex := regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)
			torExitNodeIPS = regex.FindAllString(string(body), -1)
		}
	}
	return torExitNodeIPS
}

// Validate the URI
func validURL(uri string) bool {
	_, err = url.ParseRequestURI(uri)
	return err == nil
}

// send all the request
func sendTheRequest(url string) {
	// Verify that the urls are correct.
	if validURL(url) {
		// All insecure http requests are blocked.
		if !strings.HasPrefix(url, "http://") {
			resp, err := http.Get(url)
			if err != nil {
				log.Println("Censored URL:", url)
			} else if !(url == resp.Request.URL.String()) {
				log.Println("Error URL:", url)
				wg.Add(1)
				go validateSSLCert(url)
			} else {
				fmt.Println("Valid URL:", url)
				wg.Add(1)
				go validateSSLCert(url)
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
		log.Println("Error SSL:", parsedHostname)
	} else {
		fmt.Println("Valid SSL:", parsedHostname)
	}
	wg.Done()
}
