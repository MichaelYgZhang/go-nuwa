package process

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"spider/util"
	"time"

	"golang.org/x/text/transform"
)

var (
	rateLimiter = time.Tick(
		time.Second / 10)
	verboseLogging = false
)

func SetVerboseLogging() {
	verboseLogging = true
}

func ProcessPageBody(url string) ([]byte, error) {
	<-rateLimiter
	if verboseLogging {
		log.Printf("Fetching url %s", url)
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36")
	// http.Cookie()
	// req.AddCookie()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := util.DetermineEncoding(bodyReader)
	transform.NewReader(bodyReader, e.NewEncoder())
	return ioutil.ReadAll(resp.Body)
}
