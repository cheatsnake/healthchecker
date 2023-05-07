package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/cheatsnake/healthchecker/internal/health"
)

func main() {
	var (
		urls     string
		help     bool
		urlsList []string
	)

	if len(os.Args) < 2 {
		fmt.Println("no arguments are passed \nuse -help for more info")
		os.Exit(0)
	}

	flag.Usage = printHelp
	flag.StringVar(&urls, "urls", "", "a list of HTTP servers to check")
	flag.BoolVar(&help, "help", false, "show manual page")
	flag.Parse()

	if help {
		printHelp()
	}

	if len(urls) > 0 {
		urlsList = strings.Split(urls, " ")
	}

	checker := health.NewChecker()

	s := spinner.New(spinner.CharSets[26], 200*time.Millisecond)
	s.Prefix = "Wait for responses "
	s.Start()
	results, errors := checker.Urls(urlsList)
	s.Stop()

	for i := range results {
		if i != 0 && i < len(urlsList) {
			fmt.Println("")
		}

		if errors[i] != nil {
			fmt.Printf("%d. %s\n", i+1, errors[i].Error())
			continue
		}

		fmt.Printf("%d. %s\n", i+1, results[i].ServerName)
		fmt.Println(results[i].Url)
		fmt.Println(results[i].Status)
		fmt.Println(results[i].RespondTime)

	}

}

func printHelp() {
	fmt.Println("healthcheck - show the health of HTTP servers")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("\t-urls \"...URLs\" - list of HTTP URLs to check")
	fmt.Println("\t-help - print this manual")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("\thealthcheck -urls \"https://google.com https://github.com\"")
	fmt.Println("")
	fmt.Println("Source code: https://github.com/cheatsnake/healthchecker")
	fmt.Println("Leave issue: https://github.com/cheatsnake/healthchecker/issues")
}
