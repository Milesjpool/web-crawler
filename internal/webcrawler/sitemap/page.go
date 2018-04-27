package sitemap

type Page struct {
	Url      string
	SubPages []string
}

func NewPage(url string) Page {
	return Page{Url: url, SubPages: []string{}}
}