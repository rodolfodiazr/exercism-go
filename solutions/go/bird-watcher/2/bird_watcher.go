package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var count int
    for _, b := range birdsPerDay {
        count += b
    }
    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    startDay := (week - 1) * 7
    endDay := startDay + 7
	return TotalBirdCount(
        birdsPerDay[startDay:endDay],
    )
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
        if i % 2 == 0 {
        	birdsPerDay[i]++
        }
    }
    return birdsPerDay
}
