package main

import "os"
import "github.com/Milesjpool/web-crawler/internal/webcrawler"

func main() {
  out := os.Stdout
  args := os.Args[1:]
  cliCrawler := webcrawler.CliCrawler{out}
  cliCrawler.Execute(args);
}
