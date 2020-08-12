package main

import (
	"UrlCollectGit/src"
	"flag"
	"fmt"
	"github.com/gookit/color"
	"os"
	"strconv"
)

var c src.Collector

var (
	h bool
	d string
	p int
	t string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&d, "d", "", "domain to search")
	flag.StringVar(&t, "t", "", "github token")
	flag.IntVar(&p, "p", 10, "page number")
	flag.Usage = usage
}

func usage() {
	n, _ := fmt.Fprintf(os.Stderr, `Collect urls from github
Usage: main [-d domain] [-p page] 

Options:
`)
	flag.PrintDefaults()
	_ = n
}

func start(token string) {
	color.Info.Println("[+] Start search ......")
	color.Info.Println("[+] domain: ", d, " page: ", strconv.Itoa(p))
	domain := d
	c = src.NewCollector(token)
	c.Setup(domain)
	count := 0
	for i := 1; i < p; i++ {
		results := c.Get(i, "\""+domain+"\"")
		for _, result := range results {
			count++
			color.Green.Println("[+] Num: ", count)
			color.Green.Println("   Git url: ", result.GitResult.Url)
			color.Green.Println("   Http url: ", result.UrlResult.Url)
			color.Green.Println("   Response title: ", result.UrlResult.Title)
		}
	}
	color.Info.Println("[+] Exit")
}

func main() {

	flag.Parse()
	if h {
		flag.Usage()
		return
	}

	if t == "" {
		color.Error.Println("github token is required")
		return
	}

	if d == "" {
		color.Error.Println("domain words is required")
		flag.Usage()
		return
	} else {
		start(t)
	}

}
