package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/axis-go-course/spiderman/blog"
)

func main() {
	cli := &cli{}
	flag.StringVar(&cli.templatesDir, "t", ".", "path to templates")
	flag.StringVar(&cli.bind, "bind", ":8080", "[host]:port to listen on")
	flag.Parse()

	if err := cli.run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type cli struct {
	templatesDir string
	bind         string
}

func (c *cli) run() error {
	fmt.Println("listening on", c.bind)
	s := blog.NewService(c.templatesDir)
	r := blog.NewRouter(s)
	return http.ListenAndServe(c.bind, r)
}
