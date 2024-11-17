package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Args struct {
	dirname      string
	serchcontent string
}

var args = &Args{}

func init() {
	flag.StringVar(&args.dirname, "dir", "", "-dir <dirname>")
	flag.StringVar(&args.serchcontent, "ser", "", "-ser <a word of your .look file you may look for")
}

func main() {
	flag.Parse()
	if args.dirname != "" {
		content := lookADir(fmt.Sprint("./" + args.dirname + "/.look"))
		fmt.Printf(content)
	}

	if args.serchcontent != "" {
		dirs := searchDir()
		fmt.Println(dirs)
	}
}

// location the .look file and print the content
func lookADir(lookfileLoc string) string {

	lookfileContent, err := os.ReadFile(lookfileLoc)

	if err != nil {
		fmt.Println("Read File Error: no such file")
	}

	return string(lookfileContent)
}

func searchDir() []string {
	dirS, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Read Dir Error")
	}

	results := []string{}

	for _, dir := range dirS {
		if dir.IsDir() && !strings.Contains(dir.Name(), ".") {
			lookfileLoc := fmt.Sprint("./" + dir.Name() + "/.look")
			content := lookADir(lookfileLoc)
			if strings.Contains(content, args.serchcontent) {
				results = append(results, dir.Name())
			}
		}
	}
	return results
}
