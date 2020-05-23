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
	flag.StringVar(&cli.tmplDir, "t", ".", "path to templates")
	flag.StringVar(&cli.bind, "bind", ":8080", "[host]:port to listen on")
	flag.Parse()

	if err := cli.run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type cli struct {
	tmplDir string
	bind    string
}

func (c *cli) run() error {
	fmt.Println("listening on", c.bind)
	page := blog.NewPage()
	ui, err := ui.UserInterface(c.tmplDir, page)
	if err != nil {
		return err
	}
	r := rest.NewRouter(page, ui)
	return http.ListenAndServe(c.bind, r)
}
