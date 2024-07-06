// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
package leetcode

import "fmt"

type Consecutive struct {
	start, end int
}

func (c Consecutive) isAttachable(d Consecutive) bool {
	if c.start-1 <= d.end {
		return true
	}
	if c.end+1 >= d.start {
		return true
	}

	return false
}

func (c *Consecutive) attach(d Consecutive) {
	if c.start-1 <= d.end {
		c.start = d.start
	}
	if c.end+1 >= d.start {
		c.end = d.end
	}
}

func (c Consecutive) len() int {
	return c.end - c.start + 1
}

func insert(c Consecutive, mapping map[int]Consecutive) map[int]Consecutive {
	start, end := c.start, c.end
	d, ok := mapping[start-1]
	if ok {
		d.attach(c)
		delete(mapping, start)
		return insert(d, mapping)
	}
	e, ok := mapping[end+1]
	if ok {
		e.attach(c)
		delete(mapping, end)
		return insert(e, mapping)
	}
	mapping[start] = c
	mapping[end] = c

	return mapping
}

func LongestConsecutive(nums []int) int {
	consecutiveSeq := make(map[int]Consecutive)
	for _, el := range nums {
		fmt.Println(consecutiveSeq)
		c := Consecutive{start: el, end: el}
		consecutiveSeq = insert(c, consecutiveSeq)
	}

	maxLen := 0
	for _, c := range consecutiveSeq {
		if length := c.len(); length >= maxLen {
			maxLen = length
		}
	}

	return maxLen
}
