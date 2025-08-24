//	https://leetcode.com/problems/gas-station/solutions/3011223/golang-o-n/?envType=study-plan-v2&envId=top-interview-150
//	https://leetcode.com/problems/gas-station/editorial/?envType=study-plan-v2&envId=top-interview-150

package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	fuelLeft, globalFuelLeft, start := 0, 0, 0
	for i := 0; i < n; i++ {
		globalFuelLeft += gas[i] - cost[i]
		fuelLeft += gas[i] - cost[i]
		if fuelLeft < 0 {
			start = i + 1
			fuelLeft = 0
		}
	}

	if globalFuelLeft < 0 {
		return -1
	}
	return start
}
