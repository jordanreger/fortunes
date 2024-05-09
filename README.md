# fortunes

[fortune(1)](http://man.9front.org/1/fortune) for the web.

## Usage
```go
if strings.HasPrefix(path, "/fortune") || strings.HasPrefix(path, "/fortunes") {
	fortunes.Fortune(w, r)
	return
}
```

## Contributing

Send patches/bug reports to <~jordanreger/public-inbox@lists.sr.ht>
