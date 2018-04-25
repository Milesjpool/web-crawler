package webcrawler

import "io"
import "fmt"

func Execute(writer io.Writer, args []string) {
  crawlerArgs := ParseArguments(args)
  fmt.Fprintf(writer, "Crawling '%s' from '%s'.", crawlerArgs.domain, crawlerArgs.entryPoint)
}

