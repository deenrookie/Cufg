package src

import (
	"context"
	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
	"strings"
)

func NewCollector(token string) Collector {
	return Collector{
		token: token,
	}
}

type Collector struct {
	token   string
	domain  string
	client  *github.Client
	ctx     context.Context
	Results []Result
	urlList []string
}

func (c *Collector) Setup(domain string) {
	c.ctx = context.Background()
	c.client = c.rClient()
	c.domain = domain
	c.urlList = make([]string, 0)
}

func (c *Collector) rClient() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.token},
	)
	tc := oauth2.NewClient(c.ctx, ts)
	client := github.NewClient(tc)
	return client
}

func (c *Collector) isContains(text string) bool {
	flag := 0
	for _, value := range c.urlList {
		if strings.Trim(value, "\\") == strings.Trim(text, "\\") {
			flag = 1
			break
		}
	}
	if flag == 0 {
		return false
	} else {
		return true
	}

}

func (c *Collector) Get(page int, keyword string) []Result {
	results := c.search(page, keyword)
	var tmpResults []Result
	for _, result := range results {
		status, body := Get(result.RawUrl)
		if status == "200 OK" {
			domainUrls := RegexSearch(body, c.domain)
			for _, url := range domainUrls {
				url = strings.Replace(url, " ", "", -1)
				if !c.isContains(url) {

					c.urlList = append(c.urlList, url)
					code, body := Get(url)
					titles := TitleSearch(body)
					title := ""
					if len(titles) > 0 {
						title = titles[0]
					}

					if code == "200 OK" {
						urlResult := Response{url, code, title}
						lastResult := Result{result, urlResult}
						tmpResults = append(c.Results, lastResult)
						c.Results = append(c.Results, lastResult)
					}

				}

			}
		}
	}
	return tmpResults
}

func (c *Collector) search(page int, keyword string) []SearchResult {
	listOption := github.ListOptions{
		Page: page, PerPage: 10,
	}
	searchOption := github.SearchOptions{
		Sort:        "indexed",
		Order:       "desc",
		ListOptions: listOption,
	}
	results, _, _ := c.client.Search.Code(c.ctx, keyword, &searchOption)
	rets := make([]SearchResult, 0)

	for _, result := range results.CodeResults {
		if result != nil {
			url := result.GetHTMLURL()
			repository := *result.Repository.HTMLURL
			rawUrl := TransferUrl(url)

			rets = append(rets, SearchResult{
				Repository: repository,
				Url:        url,
				RawUrl:     rawUrl,
			})

		}
	}
	return rets
}
