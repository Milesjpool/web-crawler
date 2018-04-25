package webcrawler

import (
	"io"
	"fmt"
)

type CliCrawler struct {
	Writer io.Writer
}

func (cc CliCrawler) Execute(args []string) {
	crawlerArgs := ParseCliArguments(args)
	fmt.Fprintf(cc.Writer, "Crawling '%s' from '%s'.", crawlerArgs.domain, crawlerArgs.entryPoint)

	result := Crawl(crawlerArgs)

	fmt.Fprint(cc.Writer, result.ToJson())
}
