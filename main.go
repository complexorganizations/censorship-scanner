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
		"https://4genderjustice.org",
		"https://americorps.gov",
		"https://aminoapps.com",
		"https://amp.dev",
		"https://armyfuturescommand.com",
		"https://aws.amazon.com",
		"https://azure.microsoft.com",
		"https://bandcamp.com",
		"https://bitly.com",
		"https://bloomfire.com",
		"https://bumble.com",
		"https://chaturbate.com",
		"https://cloud.google.com",
		"https://cloud.ionos.com",
		"https://cloudevents.io",
		"https://coredns.io",
		"https://cortexmetrics.io",
		"https://discord.com",
		"https://duckduckgo.com",
		"https://etcd.io",
		"https://filestage.io",
		"https://flickr.com",
		"https://fontawesome.com",
		"https://gainapp.com",
		"https://geti2p.net",
		"https://getoutline.org",
		"https://github.com",
		"https://gitlab.com",
		"https://go.dev",
		"https://goharbor.io",
		"https://golang.org",
		"https://grpc.io",
		"https://helm.sh",
		"https://hentaihaven.xxx",
		"https://imgur.com",
		"https://kubernetes.io",
		"https://mail.ru",
		"https://martus.org",
		"https://mega.io",
		"https://monday.com",
		"https://nordvpn.com",
		"https://nypost.com",
		"https://ok.ru",
		"https://opentracing.io",
		"https://operatorframework.io",
		"https://outlook.live.com",
		"https://ovhcloud.com",
		"https://prometheus.io",
		"https://redis.io",
		"https://riseup.net",
		"https://rocket.chat",
		"https://rsf.org",
		"https://ryver.com",
		"https://signal.org",
		"https://slack.com",
		"https://soundcloud.com",
		"https://sourceforge.net",
		"https://spiffe.io",
		"https://taskworld.com",
		"https://telegram.org",
		"https://thanos.io",
		"https://thepiratebay.org",
		"https://time.com",
		"https://tinder.com",
		"https://tubitv.com",
		"https://twitter.com",
		"https://ultrasurf.us",
		"https://vk.com",
		"https://vuejs.org",
		"https://weather.com",
		"https://wfrtds.org",
		"https://www.163.com",
		"https://www.360.cn",
		"https://www.academyadmissions.com",
		"https://www.accuweather.com",
		"https://www.adobe.com",
		"https://www.af.mil",
		"https://www.akamai.com",
		"https://www.alibaba.com",
		"https://www.aliexpress.com",
		"https://www.aliexpress.com",
		"https://www.amazon.com",
		"https://www.amazon.de",
		"https://www.antgroup.com",
		"https://www.apple.com",
		"https://www.atlassian.com",
		"https://www.baidu.com",
		"https://www.bbc.com",
		"https://www.bilibili.com",
		"https://www.bing.com",
		"https://www.blogger.com",
		"https://www.blogspot.com",
		"https://www.bloomberg.com",
		"https://www.booking.com",
		"https://www.businessinsider.com",
		"https://www.cachefly.com",
		"https://www.chanty.com",
		"https://www.cia.gov",
		"https://www.cisco.com",
		"https://www.cloudflare.com",
		"https://www.cncf.io",
		"https://www.cnn.com",
		"https://www.cntv.cn",
		"https://www.coccoc.com",
		"https://www.craigslist.org",
		"https://www.cricbuzz.com",
		"https://www.digitalocean.com",
		"https://www.diply.com",
		"https://www.disneyplus.com",
		"https://www.dropbox.com",
		"https://www.ea.com",
		"https://www.ebay.com",
		"https://www.ecns.cn",
		"https://www.eff.org",
		"https://www.envoyproxy.io",
		"https://www.facebook.com",
		"https://www.fandom.com",
		"https://www.fastly.com",
		"https://www.fbi.gov",
		"https://www.fc2.com",
		"https://www.federalreserveeducation.org",
		"https://www.fox.com",
		"https://www.foxnews.com",
		"https://www.freecodecamp.org",
		"https://www.fubo.tv",
		"https://www.gmw.cn",
		"https://www.google.ca",
		"https://www.google.com",
		"https://www.google.pl",
		"https://www.gumlet.com",
		"https://www.hao123.com",
		"https://www.hashicorp.com",
		"https://www.hbomax.com",
		"https://www.hostwinds.com",
		"https://www.hrw.org",
		"https://www.huffpost.com",
		"https://www.hulu.com",
		"https://www.ibm.com",
		"https://www.imdb.com",
		"https://www.instagram.com",
		"https://www.itemfix.com",
		"https://www.jaegertracing.io",
		"https://www.jd.com",
		"https://www.kanopy.com",
		"https://www.kernel.org",
		"https://www.keycdn.com",
		"https://www.kraken.com",
		"https://www.linkedin.com",
		"https://www.linode.com",
		"https://www.mediafire.com",
		"https://www.messenger.com",
		"https://www.microsoft.com",
		"https://www.msn.com",
		"https://www.nasa.gov",
		"https://www.naver.com",
		"https://www.naver.com",
		"https://www.nbc.com",
		"https://www.netflix.com",
		"https://www.netlify.com",
		"https://www.netrepid.com",
		"https://www.nicovideo.jp",
		"https://www.nsa.gov",
		"https://www.nutanix.com",
		"https://www.office.com",
		"https://www.ok.ru",
		"https://www.omct.org",
		"https://www.onet.pl",
		"https://www.openpolicyagent.org",
		"https://www.openstack.org",
		"https://www.oracle.com",
		"https://www.osce.org",
		"https://www.outbrain.com",
		"https://www.paramountplus.com",
		"https://www.paypal.com",
		"https://www.pbs.org",
		"https://www.peacocktv.com",
		"https://www.philo.com",
		"https://www.pinkcupid.com",
		"https://www.pinterest.com",
		"https://www.pixnet.net",
		"https://www.plex.tv",
		"https://www.plurk.com",
		"https://www.popads.net",
		"https://www.porn.com",
		"https://www.pornhub.com",
		"https://www.pornmd.com",
		"https://www.privateinternetaccess.com",
		"https://www.prolife.com",
		"https://www.proofhub.com",
		"https://www.proworkflow.com",
		"https://www.psiphon3.com",
		"https://www.qq.com",
		"https://www.quora.com",
		"https://www.rackspace.com",
		"https://www.reddit.com",
		"https://www.redtube.com",
		"https://www.rust-lang.org",
		"https://www.salesforce.com",
		"https://www.samsung.com",
		"https://www.sap.com",
		"https://www.servicenow.com",
		"https://www.skype.com",
		"https://www.sling.com",
		"https://www.sogou.com",
		"https://www.sohu.com",
		"https://www.soso.com",
		"https://www.spacex.com",
		"https://www.stackoverflow.com",
		"https://www.state.gov",
		"https://www.synology.com",
		"https://www.taboola.com",
		"https://www.taobao.com",
		"https://www.teamwork.com",
		"https://www.tensorflow.org",
		"https://www.tesla.com",
		"https://www.theguardian.com",
		"https://www.tianya.cn",
		"https://www.tmall.com",
		"https://www.torproject.org",
		"https://www.tribunnews.com",
		"https://www.troopmessenger.com",
		"https://www.tsa.gov",
		"https://www.tumblr.com",
		"https://www.tunnelbear.com",
		"https://www.twilio.com",
		"https://www.twitch.tv",
		"https://www.usa.gov",
		"https://www.usds.gov",
		"https://www.vmware.com",
		"https://www.vultr.com",
		"https://www.weibo.com",
		"https://www.whatsapp.com",
		"https://www.whitehouse.gov",
		"https://www.who.int",
		"https://www.wikipedia.org",
		"https://www.wordpress.com",
		"https://www.workday.com",
		"https://www.wrike.com",
		"https://www.wsj.com",
		"https://www.xswiper.com",
		"https://www.xtube.com",
		"https://www.xvideos.com",
		"https://www.yahoo.com",
		"https://www.yahoo.com",
		"https://www.youku.com",
		"https://www.youporn.com",
		"https://www.youtube.com",
		"https://www.zhihu.com",
		"https://yandex.ru",
		"https://zoom.us",
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

// Make a unique array of strings.
func makeUnique(randomStrings []string) []string {
	var uniqueString []string
	for _, value := range randomStrings {
		if !arrayContains(value, uniqueString) {
			uniqueString = append(uniqueString, value)
		}
	}
	return uniqueString
}

// Check if the array contains the value.
func arrayContains(cointains string, originalArray []string) bool {
	for _, value := range originalArray {
		if value == cointains {
			return true
		}
	}
	return false
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
func getCurrentPublicIP() string {
	var foundIP []string
	response, err := http.Get("https://api.ipengine.dev")
	handleErrors(err)
	body, err := io.ReadAll(response.Body)
	handleErrors(err)
	foundIP = regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`).FindAllString(string(body), -1)
	if len(foundIP) == 0 {
		foundIP = regexp.MustCompile(`\b(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))\b`).FindAllString(string(body), -1)
	}
	err = response.Body.Close()
	handleErrors(err)
	return foundIP[0]
}

// Obtain the public IP address of all tor exit nodes.
func getTorExitNodes() []string {
	var torExitNodeIPS []string
	response, err := http.Get("https://check.torproject.org/torbulkexitlist")
	handleErrors(err)
	body, err := io.ReadAll(response.Body)
	handleErrors(err)
	response.Body.Close()
	regex := regexp.MustCompile(`\b(?:[0-9]{1,3}\.){3}[0-9]{1,3}\b`)
	torExitNodeIPS = regex.FindAllString(string(body), -1)
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
			} else if url != resp.Request.URL.String() {
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
