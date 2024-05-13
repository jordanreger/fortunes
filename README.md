# fortunes

[fortune(1)](http://man.9front.org/1/fortune) for the web.

## Usage
```go
package main

import (
  "log"
  "net/http"

  "jordanreger.com/fortunes"
)

func main() {
  mux := http.NewServeMux()

	mux.Handle("/fortune/", fortunes.Handler())
	
  log.Fatal(http.ListenAndServe(":8080", mux))
}
```

## Contributing

Send patches/bug reports to <~jordanreger/public-inbox@lists.sr.ht>
