package util

import (
	"fmt"
	"time"
)

func TimeAgo(t time.Time) string {
	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		seconds := int(duration.Seconds())
		if seconds <= 1 {
			return "just now"
		}
		return fmt.Sprintf("%d seconds ago", seconds)
	case duration < time.Hour:
		minutes := int(duration.Minutes())
		return fmt.Sprintf("%d minutes ago", minutes)
	case duration < 24*time.Hour:
		hours := int(duration.Hours())
		return fmt.Sprintf("%d hours ago", hours)
	case duration < 30*24*time.Hour:
		days := int(duration.Hours() / 24)
		return fmt.Sprintf("%d days ago", days)
	default:
		months := int(duration.Hours() / (24 * 30))
		return fmt.Sprintf("%d months ago", months)
	}
}
