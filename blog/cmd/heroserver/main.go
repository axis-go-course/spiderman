package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/axis-go-course/spiderman/blog"
	"github.com/axis-go-course/spiderman/blog/rest"
	"github.com/axis-go-course/spiderman/blog/ui"
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
	page := blog.NewPage()
	ui := ui.UserInterface(c.templatesDir, page)
	r := rest.NewRouter(page, ui)
	return http.ListenAndServe(c.bind, r)
}
