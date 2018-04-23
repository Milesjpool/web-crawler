package main

import "os"
import "github.com/Milesjpool/web-crawler/internal/webcrawler"

func main() {
  out := os.Stdout
  webcrawler.Execute(out);
}
