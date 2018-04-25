package webcrawler

import "github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"

func Crawl(crawlerArgs *CrawlerArgs) sitemap.Page {
	return sitemap.Page{Url: crawlerArgs.entryPoint, SubPages: []sitemap.Page{}}
}
