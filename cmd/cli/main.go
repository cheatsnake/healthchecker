package main

import (
	"fmt"

	"github.com/cheatsnake/healthchecker/internal/health"
)

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://www.github.com/",
		"https://www.stackoverflow.com/",
		"https://www.medium.com/",
		"https://www.udemy.com/",
	}

	checker := health.NewChecker()

	for _, url := range urls {
		res, _ := checker.Url(url)
		fmt.Println(res.ServerName)
		fmt.Println(res.Url)
		fmt.Println(res.Status)
		fmt.Println(res.RespondTime)
	}

}
