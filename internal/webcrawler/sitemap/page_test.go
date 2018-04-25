package sitemap

import "testing"

func TestToJson(t *testing.T){
  url := "url"
  page := Page{url, []Page{}}
  actual := page.ToJson()
  expected := "{\"Url\":\""+url+"\",\"SubPages\":[]}"

  if actual != expected {
    t.Error("Expected: ", expected, ", but was: ", actual);
  }
}
