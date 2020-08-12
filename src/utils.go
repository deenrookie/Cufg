package src

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func TransferUrl(url string) string {
	url = strings.Replace(url, "https://github.com", "https://raw.githubusercontent.com/", -1)
	url = strings.Replace(url, "blob/", "", -1)
	return url
}

func RegexSearch(str string, domain string) []string {
	regex := "(?m)(http|https)://([\\w-]+\\.)+[\\w-]+(/[\\w- ./?%&=]*)?"
	var re = regexp.MustCompile(regex)
	urls := make([]string, 0)
	for _, match := range re.FindAllString(str, -1) {
		if strings.Contains(match, domain) {
			urls = append(urls, match)
		}
	}
	return urls
}

func TitleSearch(str string) []string {
	re := regexp.MustCompile(`(?m)<title>([\s\S]*?)</title>`)
	titles := make([]string, 0)
	for _, match := range re.FindAllString(str, -1) {
		titles = append(titles, match)
	}
	return titles
}

func Get(url string) (code string, body string) {
	client := http.Client{Timeout: 1 * time.Second}
	request, err := http.NewRequest("GET", url, nil)
	if err == nil {
		response, err := client.Do(request)
		if err != nil {
			return "", ""
		}
		tmp, _ := ioutil.ReadAll(response.Body)
		body = string(tmp)
		return response.Status, body
	}
	return "", ""
}
