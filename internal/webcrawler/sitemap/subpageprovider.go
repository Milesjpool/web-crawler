package sitemap

type SubPageProvider interface {
	GetSubPages(url string) []string
}