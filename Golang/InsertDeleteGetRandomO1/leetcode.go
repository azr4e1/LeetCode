// Implement the RandomizedSet class:

// RandomizedSet() Initializes the RandomizedSet object.
// bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
// bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
// int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
// You must implement the functions of the class such that each function works in average O(1) time complexity.
package leetcode

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	indexValue   map[int]int
	valueIndex   map[int]int
	highestIndex int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indexValue: make(map[int]int),
		valueIndex: make(map[int]int),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.valueIndex[val]; ok {
		return false
	}
	rs.highestIndex++
	rs.indexValue[rs.highestIndex] = val
	rs.valueIndex[val] = rs.highestIndex
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	index, ok := rs.valueIndex[val]
	if !ok {
		return false
	}
	delete(rs.valueIndex, val)
	delete(rs.indexValue, index)
	if rs.highestIndex > 1 {
		highestEl := rs.indexValue[rs.highestIndex]
		rs.valueIndex[highestEl] = index
		rs.indexValue[index] = highestEl
	}
	rs.highestIndex--
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	randomInd := rand.Intn(rs.highestIndex) + 1
	return rs.indexValue[randomInd]
}

func (rs RandomizedSet) String() string {
	var output string = "[ "
	for el := range rs.valueIndex {
		output += fmt.Sprintf("%d ", el)
	}
	return output + "]"
}
