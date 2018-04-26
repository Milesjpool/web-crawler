package sitemap

type SubPageProvider interface {
	GetSubPages(page string) []string
}