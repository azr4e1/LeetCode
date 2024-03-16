// There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

// You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

// Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique
package leetcode

func CanCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	distances := make([]int, n)
	var total int
	for i := 0; i < n; i++ {
		distance := gas[i] - cost[i]
		distances[i] = distance
		total += distance
	}
	if total < 0 {
		return -1
	}
	index := 0
	startIndex := 0
	partialSum := 0
	for {
		partialSum += distances[index]
		index = (index + 1) % n
		if partialSum < 0 {
			partialSum = 0
			startIndex = index
			continue
		}
		if (index+1)%n == startIndex {
			return startIndex
		}
	}
}
