package request

import (
	"fmt"
	"io/ioutil"
	. "music163-hotcomments-spider/util"
	"net/http"
	"strings"
)

var client = &http.Client{}

func post(url string, data string) string {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	Check(err)
	req.Header = getHeader()
	return getBody(client.Do(req))
}

func get(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	Check(err)
	req.Header = getHeader()
	return getBody(client.Do(req))
}

func getBody(resp *http.Response, err error) string {
	Check(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		Check(fmt.Errorf("statuc code error : %v", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	Check(err)

	return ToString(body)
}
