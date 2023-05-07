package health

import "time"

const requestTimeout = 5 * time.Second

const titleRgPattern string = `<title.*>(.*?)</title.*>`

const (
	successCodesThreshold  = 300
	redirectCodesThreshold = 400
)

const (
	fastSpeedThreshold   = 500
	mediumSpeedThreshold = 1000
)
