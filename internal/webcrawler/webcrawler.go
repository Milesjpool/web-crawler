package webcrawler

import (
	"github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/subpageproviders"
)

type WebCrawler struct {
	SubPageProvider subpageproviders.SubPageProvider
}

func (wc WebCrawler) Crawl(crawlerArgs CrawlerArgs) sitemap.PageListing {
	pageListing := sitemap.NewPageListing(crawlerArgs.GetDomain(), crawlerArgs.GetEntryPoint())

	entryPoint := pageListing.GetEntryPoint()
	subPages := wc.SubPageProvider.GetSubPages(entryPoint)
	for _, url := range subPages {
		pageListing.Upsert(entryPoint.String(), url.String())
	}
	return pageListing
}