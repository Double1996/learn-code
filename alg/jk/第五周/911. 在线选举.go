package 第五周

import "sort"

type TopVotedCandidate struct {
	persons []int
	times   []int
	result  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	sta := map[int]int{}
	result := make([]int, len(times))
	sta[persons[0]] = 1
	result[0] = persons[0]
	for i := 1; i < len(persons); i++ {
		lastTimes, _ := sta[result[i-1]]
		maxTimes := lastTimes

		times, _ := sta[persons[i]]
		if times+1 >= maxTimes {
			result[i] = persons[i]
		} else {
			result[i] = result[i-1]
		}

		sta[persons[i]] += 1

	}

	return TopVotedCandidate{
		persons: persons,
		times:   times,
		result:  result,
	}

}

func (this *TopVotedCandidate) Q(t int) int {
	index := sort.SearchInts(this.times, t)
	if index == len(this.result) {
		index--
	} else {
		if this.times[index] != t {
			index--
		}
	}

	return this.result[index]
}
