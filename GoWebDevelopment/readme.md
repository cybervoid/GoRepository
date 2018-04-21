# A section for getting started with Go web development.
This workspace is for learning Go web development.

## Setting up environment
1. Set up a working Go environment (not covered in this workspace).
2. Install GopherJS: `go get -u github.com/gopherjs/gopherjs`


### Using GopherJS
1. Useful commands
  1. Build the JS Source file to be included in the web page: `$gopherjs build`
  2. Build and minify the JS source file: `$gopherjs build -m`
  3. Compile Go packages and start the HTTP server: `$gopherjs serve`
2. Performance tips:
  1. Use the -m command line flag to generate minified
  2. Enable gzip compression on your web server
  3. Use int instead of (u)int8/16/32/64
  4. Use float64 instead of float32

## Workspace content
1. [Gopherface](./Gopherface) Social Media Web Application, based upon course content in: [LEARNING PATH: Go: Advancing into Web Development with Go](https://www.udemy.com/learning-path-go-advancing-into-web-development-with-go/)
  1. Source Code: `go get github.com/EngineerKamesh/gofullstack`
