package main

import "os"
import "github.com/Milesjpool/web-crawler/internal/webcrawler"

func main() {
	args := os.Args[1:]
	cliCrawler := webcrawler.CliCrawler{Writer: os.Stdout}
	cliCrawler.Execute(args);
}
