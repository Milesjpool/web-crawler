package main

import "fmt"
import "github.com/Milesjpool/web-crawler/internal/webcrawler"

func main() {
  output := webcrawler.Crawl();
  fmt.Println(output)
}
