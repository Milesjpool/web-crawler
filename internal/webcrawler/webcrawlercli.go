package webcrawler

import (
  "io"
  "fmt"
  "github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"
)

func Execute(writer io.Writer, args []string) {
  crawlerArgs := ParseCliArguments(args)
  fmt.Fprintf(writer, "Crawling '%s' from '%s'.", crawlerArgs.domain, crawlerArgs.entryPoint)
  result := Crawl(crawlerArgs)
  fmt.Fprint(writer, result.ToJson())
}


func Crawl(crawlerArgs *CrawlerArgs) sitemap.Page {
  return sitemap.Page{Url: crawlerArgs.entryPoint, SubPages: []sitemap.Page{}}
}

