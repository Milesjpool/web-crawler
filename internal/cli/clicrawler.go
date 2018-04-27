package cli

import (
	"io"
	"fmt"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/subpageproviders"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/http"
	"github.com/Milesjpool/web-crawler/internal/webcrawler"
)

type CliCrawler struct {
	Writer io.Writer
	HttpClient http.HttpClient
}

func (cc CliCrawler) Execute(args []string) {
	crawlerArgs := ParseCliArguments(args)
	crawler := InitialiseWebCrawler(cc)
	result := crawler.Crawl(crawlerArgs)
	fmt.Fprint(cc.Writer, result.ToJson())
}

func InitialiseWebCrawler(cc CliCrawler) webcrawler.WebCrawler {
	pageRetriever := http.WebPageRetriever{HttpClient: cc.HttpClient}
	htmlSubPageProvider := subpageproviders.HtmlSubPageProvider{PageRetriever: pageRetriever}
	absoluteSubPageProvider := subpageproviders.AbsoluteSubPageProvider{SubPageProvider: htmlSubPageProvider}
	domainSubPageProvider := subpageproviders.DomainSubPageProvider{SubPageProvider: absoluteSubPageProvider}

	return webcrawler.WebCrawler{SubPageProvider: domainSubPageProvider}
}
