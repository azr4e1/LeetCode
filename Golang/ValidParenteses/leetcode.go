package leetcode

import "errors"

var parentheses = map[byte]byte{
	'(': ')',
	'{': '}',
	'[': ']',
}

type Queue []byte

func (q *Queue) Add(el byte) {
	*q = append(*q, el)
}

func (q *Queue) Pop() (byte, error) {
	var el byte
	if len(*q) == 0 {
		return el, errors.New("Queue is empty")
	}
	*q, el = (*q)[:len(*q)-1], (*q)[len(*q)-1]

	return el, nil
}

func NewQueue() *Queue {
	return &Queue{}
}

func IsValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	q := NewQueue()
	for i := 0; i < len(s); i++ {
		el := s[i]
		val, ok := parentheses[el]
		if ok {
			q.Add(val)
			continue
		}
		closing, err := q.Pop()
		if err != nil || closing != el {
			return false
		}
	}
	return len(*q) == 0
}
