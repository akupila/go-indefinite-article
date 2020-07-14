package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/akupila/go-indefinite-article"
)

func main() {
	only := flag.Bool("only", false, "Only output article")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <word>\n", os.Args[0])
		os.Exit(2)
	}

	word := os.Args[1]
	article := indefinite.Article(word)
	if *only {
		fmt.Fprint(os.Stdout, article)
		return
	}
	fmt.Fprintf(os.Stdout, "%s %s\n", indefinite.Article(word), word)
	return
}
