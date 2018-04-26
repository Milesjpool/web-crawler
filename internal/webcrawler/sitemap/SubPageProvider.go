package sitemap

type SubPageProvider interface {
	GetSubPages(page string) []string
}

type HtmlSubPageProvider struct {
}

func (s HtmlSubPageProvider) GetSubPages(page string) []string {
	return []string{}
}