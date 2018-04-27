package main

import (
	"os"
	"net/http"
	"github.com/Milesjpool/web-crawler/internal/cli"
)

func main() {
	args := os.Args[1:]
	cliCrawler := cli.CliCrawler{Writer: os.Stdout, HttpClient:http.DefaultClient}
	cliCrawler.Execute(args);
}
