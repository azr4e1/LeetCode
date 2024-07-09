// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
package leetcode

import "fmt"

type Chain struct {
	start, end int
}

func (c *Chain) String() string {
	return fmt.Sprintf("[%d %d]", c.start, c.end)
}

func (c *Chain) Concat(d Chain) {
	switch {
	case d.start <= c.start && d.end >= c.end:
		c.start = d.start
		c.end = d.end
	case d.start >= c.start && d.end <= c.end:
		// do nothing
	case d.start <= c.start && d.end >= c.start-1:
		c.start = d.start
	case d.start <= c.end+1 && d.end >= c.end:
		c.end = d.end
	}
}

func (c Chain) Len() int {
	return c.end - c.start + 1
}

func LongestConsecutive(nums []int) int {
	var seen = make(map[int]*Chain)
	for _, e := range nums {
		chain := Chain{start: e, end: e}
		seen[e] = &chain
	}

	maxLen := 0
	for _, v := range seen {
		el := v.start - 1
		c, ok := seen[el]
		if ok {
			v.Concat(*c)
			seen[el] = v
		}
		el = v.end + 1
		c, ok = seen[el]
		if ok {
			v.Concat(*c)
			seen[el] = v
		}
		if newLen := v.Len(); newLen > maxLen {
			maxLen = newLen
		}
	}
	return maxLen
}
