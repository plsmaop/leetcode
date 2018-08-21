/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
import "sort"

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	} else if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return intervals[i].Start < intervals[j].Start
	})
	p1, end := intervals[0].Start, intervals[0].End

	ans := []Interval{}
	for _, interval := range intervals {
		if p1 == interval.Start {
			end = interval.End
			continue
		}
		if interval.Start <= end {
			if interval.End > end {
				end = interval.End
			}
			continue
		}
		ans = append(ans, Interval{p1, end})
		p1 = interval.Start
		end = interval.End
	}
	ans = append(ans, Interval{p1, end})
	return ans
}