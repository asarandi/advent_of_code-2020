package main

import (
	"container/ring"
	"fmt"
)

func main() {
	const (
		iter = 10000000
		max  = 1000000
	)
	//	nums := []int{3, 8, 9, 1, 2, 5, 4, 6, 7} // sample
	nums := []int{5, 8, 3, 9, 7, 6, 2, 4, 1}
	for i := 10; i <= max; i++ {
		nums = append(nums, i)
	}
	ptr := map[int]*ring.Ring{}
	n := len(nums)
	r := ring.New(n)
	for i := 0; i < n; i++ {
		r.Value = nums[i]
		ptr[nums[i]] = r
		r = r.Next()
	}
	for i := 0; i < iter; i++ {
		three := r.Unlink(3)
		vals := map[int]bool{}
		for j := 0; j < 3; j++ {
			vals[three.Value.(int)] = true
			three = three.Next()
		}
		dv := r.Value.(int) - 1
		for ; vals[dv]; dv-- {
		}
		if dv < 1 {
			dv = max
		}
		for ; vals[dv]; dv-- {
		}
		ptr[dv].Link(three)
		r = r.Next()
	}
	a := ptr[1].Next().Value.(int)
	b := ptr[1].Next().Next().Value.(int)
	fmt.Println("part 2:", a*b)
}
