package webcrawler

import (
	"io"
	"fmt"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"
)

type CliCrawler struct {
	Writer io.Writer
}

func (cc CliCrawler) Execute(args []string) {
	crawlerArgs := ParseCliArguments(args)
	fmt.Fprintf(cc.Writer, "Crawling '%s' from '%s'.", crawlerArgs.domain, crawlerArgs.entryPoint)

	subPageProvider := sitemap.HtmlSubPageProvider{}
	result := WebCrawler{SubPageProvider: subPageProvider}.Crawl(crawlerArgs)

	fmt.Fprint(cc.Writer, result.ToJson())
}
