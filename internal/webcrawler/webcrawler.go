package webcrawler

import (
	"github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/subpageproviders"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/util"
)

type WebCrawler struct {
	SubPageProvider subpageproviders.SubPageProvider
}

func (wc WebCrawler) Crawl(crawlerArgs CrawlerArgs) sitemap.PageListing {
	pageListing := sitemap.NewPageListing(crawlerArgs.GetDomain(), crawlerArgs.GetEntryPoint())

	for i := 0; i<len(pageListing.Pages); i++ {
		page := pageListing.Pages[i]

		subpages := wc.SubPageProvider.GetSubPages(page.GetUrl())
		pageListing.AddSubPages(page.Url, util.AsStrings(subpages))
	}

	return pageListing
}
