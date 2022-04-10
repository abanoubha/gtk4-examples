package main

import (
	"fmt"

	"github.com/abanoubha/gtk"
)

func main() {
	what := gtk.UseGtk()
	fmt.Println("using gtk status :", what)
}
