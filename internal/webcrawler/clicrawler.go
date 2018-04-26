package webcrawler

import (
	"io"
	"fmt"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/http"
)

type CliCrawler struct {
	Writer io.Writer
}


func (cc CliCrawler) Execute(args []string) {
	crawlerArgs := ParseCliArguments(args)
	fmt.Fprintf(cc.Writer, "Crawling '%s' from '%s'.", crawlerArgs.domain, crawlerArgs.entryPoint)

	pageRetriever := http.WebPageRetriever{}
	subPageProvider := sitemap.HtmlSubPageProvider{PageRetriever: pageRetriever}
	crawler := WebCrawler{SubPageProvider: subPageProvider}

	result := crawler.Crawl(crawlerArgs)

	fmt.Fprint(cc.Writer, result.ToJson())
}
