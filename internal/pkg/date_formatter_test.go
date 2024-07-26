package pkg

import (
    "testing"
    "time"
)

// TestFormatDate tests the FormatDate function
func TestFormatDate(t *testing.T) {
    tests := []struct {
        input    time.Time
        expected string
    }{
        {time.Date(2024, time.January, 1, 15, 4, 0, 0, time.UTC), "January 1st, 2024 at 3:04 PM"},
        {time.Date(2024, time.January, 2, 15, 4, 0, 0, time.UTC), "January 2nd, 2024 at 3:04 PM"},
        {time.Date(2024, time.January, 3, 15, 4, 0, 0, time.UTC), "January 3rd, 2024 at 3:04 PM"},
        {time.Date(2024, time.January, 4, 15, 4, 0, 0, time.UTC), "January 4th, 2024 at 3:04 PM"},
        {time.Date(2024, time.January, 11, 15, 4, 0, 0, time.UTC), "January 11th, 2024 at 3:04 PM"},
        {time.Date(2024, time.January, 21, 15, 4, 0, 0, time.UTC), "January 21st, 2024 at 3:04 PM"},
        {time.Date(2024, time.January, 22, 15, 4, 0, 0, time.UTC), "January 22nd, 2024 at 3:04 PM"},
        {time.Date(2024, time.January, 23, 15, 4, 0, 0, time.UTC), "January 23rd, 2024 at 3:04 PM"},
    }

    for _, tt := range tests {
        t.Run(tt.expected, func(t *testing.T) {
            result := FormatDate(tt.input)
            if result != tt.expected {
                t.Errorf("FormatDate(%v) = %v; want %v", tt.input, result, tt.expected)
            }
        })
    }
}

// TestGetDaySuffix tests the getDaySuffix function
func TestGetDaySuffix(t *testing.T) {
    tests := []struct {
        input    int
        expected string
    }{
        {1, "st"},
        {2, "nd"},
        {3, "rd"},
        {4, "th"},
        {11, "th"},
        {12, "th"},
        {13, "th"},
        {21, "st"},
        {22, "nd"},
        {23, "rd"},
        {24, "th"},
    }

    for _, tt := range tests {
        t.Run(tt.expected, func(t *testing.T) {
            result := getDaySuffix(tt.input)
            if result != tt.expected {
                t.Errorf("getDaySuffix(%v) = %v; want %v", tt.input, result, tt.expected)
            }
        })
    }
}
