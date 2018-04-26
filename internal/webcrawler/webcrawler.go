package webcrawler

import "github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"

type WebCrawler struct {
	SubPageProvider sitemap.SubPageProvider
}

func (wc WebCrawler) Crawl(crawlerArgs *CrawlerArgs) sitemap.Page {
	subPages := wc.SubPageProvider.GetSubPages("")
	root := sitemap.NewPage(crawlerArgs.entryPoint)
	for _, url := range subPages {
		root.AddSubPage(url)
	}
	return root
}