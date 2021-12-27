package contest268

/*
 https://leetcode.com/contest/weekly-contest-268/problems/watering-plants/
 */
func wateringPlants(plants []int, capacity int) int {

	var steps int
	full := capacity
	for i:=0;i<len(plants);i++ {
		if capacity >= plants[i] {
			steps ++
			capacity = capacity - plants[i]
		} else {
			steps = steps + 2*(i+1)-1
			capacity = full
			capacity = capacity -  plants[i]
		}
	}
	return steps
}
