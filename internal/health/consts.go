package health

import "time"

const requestTimeout = 7 * time.Second

const (
	titleRgPattern string = `<title.*>(.*?)</title.*>`
	digitRgPattern string = `[0-9]+`
)

const (
	successCodesThreshold  = 300
	redirectCodesThreshold = 400
)

const (
	fastSpeedThreshold   = 500
	mediumSpeedThreshold = 1000
)
