package cli

import "github.com/Milesjpool/web-crawler/internal/webcrawler/domain"

type CliCrawlerArgs struct {
	entryPoint string
	domain     string
}

func (s CliCrawlerArgs) GetDomain() string  {
	return s.domain
}

func (s CliCrawlerArgs) GetEntryPoint() string  {
	return s.entryPoint
}

func ParseCliArguments(rawArgs []string) (args *CliCrawlerArgs) {
	url := rawArgs[0]
	host, _ := domain.From(url)

	return &CliCrawlerArgs{entryPoint: url, domain: host}
}
