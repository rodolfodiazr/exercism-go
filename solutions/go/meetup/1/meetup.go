package meetup

import (
    "time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

const (
    First WeekSchedule = iota
    Second
    Third
    Fourth
    Last
    Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
    var startDay, endDay int
    switch wSched {
    case First:
        startDay, endDay = 1, 7
    case Second:
        startDay, endDay = 8, 14
    case Third:
        startDay, endDay = 15, 21
    case Fourth:
        startDay, endDay = 22, 28
    case Teenth:
        startDay, endDay = 13, 19
    case Last:
        date := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
        for {
            if date.Weekday() == wDay {
                return date.Day()
            }
            date = date.AddDate(0, 0, -1)
        }
    }

    for day := startDay; day <= endDay; day++ {
        date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
        if date.Weekday() == wDay {
            return day
        }
    }

    return -1
}
