package main

import (
	"fmt"
	"sort"
)

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.



Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.


Constraints:

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 3}}))
	fmt.Println(merge([][]int{{}}))
	fmt.Println(merge([][]int{{1, 4}, {5, 6}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 1}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
	fmt.Println(merge([][]int{{0, 0}, {0, 0}, {4, 4}, {0, 0}, {1, 3}, {5, 5}, {4, 6}, {1, 1}, {0, 2}}))
	fmt.Println(merge([][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}))
}

func merge(intervals [][]int) [][]int {
	var res [][]int

	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastEnd := res[len(res)-1][1]

		if intervals[i][0] <= lastEnd {
			res[len(res)-1][1] = max(lastEnd, intervals[i][1])
		} else {
			res = append(res, []int{intervals[i][0], intervals[i][1]})
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
