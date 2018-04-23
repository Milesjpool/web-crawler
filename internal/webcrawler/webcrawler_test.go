package webcrawler

import "testing"

func TestHello(t *testing.T){
  if Crawl() != "Hello, world!" {
    t.Error("incorrect greeting");
  }
}
