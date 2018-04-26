package main

import "os"
import (
	"github.com/Milesjpool/web-crawler/internal/webcrawler"
	"net/http"
)

func main() {
	args := os.Args[1:]
	cliCrawler := webcrawler.CliCrawler{Writer: os.Stdout, HttpClient:http.DefaultClient}
	cliCrawler.Execute(args);
}
