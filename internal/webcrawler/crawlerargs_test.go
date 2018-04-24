package webcrawler

import "testing"

func TestParseArguments(t *testing.T){
  args := []string{"http://my.domain.com/path"}
  crawlerArgs := ParseArguments(args)

  expectedDomain := "my.domain.com"

  if crawlerArgs.domain != expectedDomain {
    t.Error("Expected: ", expectedDomain, ", but was: ", crawlerArgs.domain);
  }
}
