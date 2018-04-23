package webcrawler

import "testing"

func TestExecute(t *testing.T){
  writer := new(MockWriter)

  Execute(writer)

  actual := writer.written
  expected := "Hello, world!"

  if actual != expected {
    t.Error("Expected: ", expected, ", but was: ", actual);
  }
}


type MockWriter struct {
  written string
}

func (mw *MockWriter) Write(b []byte) (int, error) {
  mw.written = string(b[:])
  return 1, nil
}
