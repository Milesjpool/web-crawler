package webcrawler

type CrawlerArgs interface {
	GetEntryPoint() string
	GetDomain()     string
}