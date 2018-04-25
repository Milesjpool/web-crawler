package webcrawler

import "testing"

func TestExecute(t *testing.T){
  writer := new(MockWriter)
  page := "http://my.domain.com/path"
  args := []string{page}

  crawler := CliCrawler{writer}
  crawler.Execute(args)

  actual := writer.written
  expected := []string{
    "Crawling 'my.domain.com' from '" + page + "'.",
    "{\"Url\":\"" + page + "\",\"SubPages\":[]}",
  }

  if len(actual) != len(expected) {
    t.Error("Incorrect output - Expected: ", expected, ", but was: ", actual)
    return
  }
  for i, v := range expected {
    if v != actual[i] {
      t.Error("Expected: ", v, ", but was: ", actual[i])
    }
  }
}


type MockWriter struct {
  written []string
}

func (mw *MockWriter) Write(b []byte) (int, error) {
  mw.written = append(mw.written, string(b[:]))
  return 1, nil
}