package date

import "time"

func Today() string {
	currentTime := time.Now()

	return currentTime.Format("2006-01-02")
}
