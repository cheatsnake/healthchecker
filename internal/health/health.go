package health

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Result struct {
	ServerName  string
	Url         string
	Status      string
	RespondTime string
}

type Checker struct {
	client *http.Client
}

func NewChecker() *Checker {
	client := &http.Client{
		Timeout: requestTimeout,
	}

	return &Checker{
		client: client,
	}
}

func (ch *Checker) Url(url string) (Result, error) {
	start := time.Now()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Result{}, fmt.Errorf("passed URL is not valid: %s", url)
	}

	res, err := ch.client.Do(req)
	if err != nil {
		return Result{}, fmt.Errorf("failed to complete the request at the specified URL: %s", url)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Result{}, fmt.Errorf("failed to read the response from the server: %s", url)
	}

	finish := time.Since(start).Milliseconds()

	return Result{
		Url:         url,
		ServerName:  retriveServerName(body),
		Status:      defineStatus(res.StatusCode, res.Status),
		RespondTime: defineSpeed(finish),
	}, nil
}

func (ch *Checker) Urls(urls []string) ([]Result, []error) {
	results := make([]Result, 0, len(urls))
	errors := make([]error, 0, len(urls))
	done := make(chan bool)

	for _, url := range urls {
		go func(u string) {
			res, err := ch.Url(u)

			results = append(results, res)
			errors = append(errors, err)

			done <- true
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		<-done
	}

	return results, errors
}
