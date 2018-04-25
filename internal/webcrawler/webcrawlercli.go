package webcrawler

import (
  "io"
  "fmt"
)

func Execute(writer io.Writer, args []string) {
  crawlerArgs := ParseCliArguments(args)
  fmt.Fprintf(writer, "Crawling '%s' from '%s'.", crawlerArgs.domain, crawlerArgs.entryPoint)
  result := "{\"page\":\"" + crawlerArgs.entryPoint + "\", \"subpages\":[]}"
  fmt.Fprint(writer, result)
}

