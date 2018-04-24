package webcrawler

import "github.com/Milesjpool/web-crawler/internal/webcrawler/domain"

type CrawlerArgs struct {
  domain string
}

func ParseArguments(rawArgs []string) (args *CrawlerArgs){
  domain, _ := domain.ExtractDomain(rawArgs[0])

  return &CrawlerArgs{domain: domain}
}
