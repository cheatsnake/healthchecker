package health

import (
	"fmt"
	"io"
	"net/http"
	"sync"
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

func (ch *Checker) url(url string) (Result, error) {
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

type multiResult struct {
	results []Result
	errors  []error
	mu      sync.Mutex
}

func (ch *Checker) Urls(urls []string) ([]Result, []error) {
	var wg sync.WaitGroup
	mr := multiResult{
		results: make([]Result, len(urls)),
		errors:  make([]error, len(urls)),
	}

	wg.Add(len(urls))

	for idx, url := range urls {
		go func(i int, u string) {
			defer wg.Done()

			res, err := ch.url(u)

			mr.mu.Lock()
			mr.results[i] = res
			mr.errors[i] = err
			mr.mu.Unlock()
		}(idx, url)
	}

	wg.Wait()

	return mr.results, mr.errors
}
