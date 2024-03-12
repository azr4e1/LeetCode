package leetcode

import "fmt"

type Vote struct {
	Candidate int
	Vote      int
}

func MajorityElement(nums []int) int {
	vote := new(Vote)
	for _, v := range nums {
		fmt.Println(vote)
		if vote.Vote == 0 {
			vote.Candidate = v
			vote.Vote = 1
			continue
		}
		if v == vote.Candidate {
			vote.Vote++
		} else {
			vote.Vote--
		}
	}

	return vote.Candidate
}
