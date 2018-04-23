package webcrawler

import "testing"

func TestExecute(t *testing.T){
  writer := new(MockWriter)
  args := []string{"planet"}

  Execute(writer, args)

  actual := writer.written
  expected := "Hello, planet!"

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
