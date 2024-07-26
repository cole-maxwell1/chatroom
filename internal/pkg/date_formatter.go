package pkg

import "time"

// formatDate formats the given time.Time into the desired format
func FormatDate(t time.Time) string {
    day := t.Day()
    daySuffix := getDaySuffix(day)
    formattedDate := t.Format("January 2") + daySuffix + t.Format(", 2006 at 3:04 PM")
    return formattedDate
}

// getDaySuffix returns the suffix for the day of the month
func getDaySuffix(day int) string {
    if day >= 11 && day <= 13 {
        return "th"
    }
    switch day % 10 {
    case 1:
        return "st"
    case 2:
        return "nd"
    case 3:
        return "rd"
    default:
        return "th"
    }
}