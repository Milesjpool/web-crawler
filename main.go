package main

import "os"
import "github.com/Milesjpool/web-crawler/internal/webcrawler"

func main() {
  out := os.Stdout
  args := os.Args[1:]
  webcrawler.Execute(out, args);
}
