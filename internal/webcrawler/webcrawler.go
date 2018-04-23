package webcrawler

import "io"
import "fmt"

func Execute(writer io.Writer, args []string) {
  arg := args[0]
  fmt.Fprintf(writer, "Hello, %s!", arg)
}
