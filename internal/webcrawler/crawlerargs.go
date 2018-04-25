package webcrawler

import "github.com/Milesjpool/web-crawler/internal/webcrawler/domain"

type CrawlerArgs struct {
  entryPoint string
  domain string
}

func ParseArguments(rawArgs []string) (args *CrawlerArgs){
  url := rawArgs[0]
  host, _ := domain.From(url)

  return &CrawlerArgs{entryPoint: url, domain: host}
}
