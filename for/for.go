package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var sum int
	for i := 0; i < len(birdsPerDay); i++{
        sum+= birdsPerDay[i]
    }
return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var sum int
    if week % 2 == 0 {
		for i := 7; i < 14; i++{			
			sum+= birdsPerDay[i]
		}
	return sum
    }
	
	for i := 0; i < 7 ; i++{
	sum+= birdsPerDay[i]
		}
	
	return sum
}
// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay) ; i++{
		if i % 2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] +1
		}
	}

return birdsPerDay
}
