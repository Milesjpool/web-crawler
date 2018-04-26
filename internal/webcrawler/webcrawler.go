package webcrawler

import "github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"

type WebCrawler struct {
	SubPageProvider sitemap.SubPageProvider
}

func (wc WebCrawler) Crawl(crawlerArgs *CrawlerArgs) sitemap.Page {
	root := sitemap.NewPage(crawlerArgs.entryPoint)

	subPages := wc.SubPageProvider.GetSubPages(root.GetUrl())
	for _, url := range subPages {
		page := sitemap.NewPage(url.String())
		root.AddSubPage(page)
	}
	return root
}