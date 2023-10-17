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
	for i := len(birdsPerDay)/2 - 1; i < len(birdsPerDay); i++{			
        sum+= birdsPerDay[i]
    }
	return sum
    }
	
    	for i := 0; i < len(birdsPerDay)/2 - 1; i++{
        sum+= birdsPerDay[i]
   		 }
      
	  return sum
}
// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	newS := birdsPerDay
	newS[0]=newS[0]+1
	return newS
}
