package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/AldieNightStar/minforth/minforth"
	"golang.design/x/clipboard"
)

func main() {
	// Open the file app.forth
	f, err := os.Open("app.forth")
	if err != nil {
		log.Fatal(err)
	}
	fdat, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	src := string(fdat)

	// Compile the file
	compiled, err := minforth.Compile("cell1", "message1", src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(compiled)

	// Paste into the clipboard
	err = clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}

	clipboard.Write(clipboard.FmtText, []byte(compiled))
}
