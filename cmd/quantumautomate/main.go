// cmd/quantumautomate/main.go
package main

import (
"flag"
"log"
"os"

"quantumautomate/internal/quantumautomate"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := quantumautomate.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
