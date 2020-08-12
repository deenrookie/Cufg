package src

type SearchResult struct {
	Repository string
	Url        string
	RawUrl     string
}

type Response struct {
	Url string
	Status string
	Title  string
}

type Result struct {
	GitResult SearchResult
	UrlResult Response
}
