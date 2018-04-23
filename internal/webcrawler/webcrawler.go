package webcrawler

import "io"
import "fmt"

func Execute(writer io.Writer) {
  fmt.Fprintf(writer, "Hello, world!")
}
