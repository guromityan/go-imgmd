package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/guromityan/go-imgmd/lib"
	"gopkg.in/alecthomas/kingpin.v2"
)

const version = "1.0.0"

var (
	app = kingpin.New("imgmd", "Convert image to Markdown.")

	target = app.Arg("target", "Target directory containing images.").Required().ExistingDir()
	output = app.Flag("output", "Markdown file name to output.").Short('o').Default("output.md").String()
)

func main() {
	app.Version(version)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	stmts, err := lib.Dirwalk(*target)
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create(*output)
	if err != nil {
		log.Fatalf("Failed to create %v: %v", output, err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range stmts {
		fmt.Fprintln(w, line)
	}
	err = w.Flush()
	if err != nil {
		log.Fatalf("Failed to write file %v: %v", output, err)
	}
}
