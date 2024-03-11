package date

import (
	"fmt"
	"time"
)

// Format formats the given date using the given format
func Format(date time.Time, format string) (string, error) {
	switch format {
	case "YYYY-MM-DD":
		return date.Format("2006-01-02"), nil
	case "MM-DD-YYYY":
		return date.Format("01-02-2006"), nil
	}
	return "", fmt.Errorf("unsupported format type: %s", format)
}
