package main

import (
	"fmt"
	"sort"
)

// implements Binary Index Tree (Fenwick tree)
// https://en.wikipedia.org/wiki/Fenwick_tree

func countSmaller(nums []int) []int {
	nlen := len(nums)

	sorted := make([]int, nlen)
	copy(sorted, nums)
	sort.Ints(sorted)

	mp := make(map[int]int)
	index := 0
	for i := range sorted {
		if _, ok := mp[sorted[i]]; !ok {
			mp[sorted[i]] = index
			index++
		}
	}

	// fmt.Println(sorted)
	// fmt.Println(mp)
	// fmt.Println(nums)

	bit := make(BIT, len(mp)+1)

	res := make([]int, nlen)
	// loop over original input backwards
	for i := nlen - 1; i >= 0; i-- {
		fmt.Printf("***** number: %v index -1: %v\n", nums[i], mp[nums[i]]-1)
		// sum(index - 1)
		res[i] = sum(bit, mp[nums[i]]-1)
		add(bit, mp[nums[i]], 1)
	}
	return res
}

type BIT []int

// builds the BIT, i increases with the opp i += i & (-i)
// and it increases it increments by the value
func add(b BIT, index int, value int) {
	//  fmt.Printf("init index %v b%v...\n", index, b)
	i := index + 1
	for i < len(b) {
		b[i] += value
		i += i & (-i)
		// fmt.Printf("curr b %v i %v value %v\n", b, i, value)
	}
	// fmt.Println(".")
}

// sum up going backwards through the binary index tree
// they index, key, is found with i = i - i&(-i)
// an opp that finds the last 1 bit, and flips it.
func sum(b BIT, index int) int {
	sum := 0
	//fmt.Printf("init index %v b%v...\n", index, b)
	for i := index + 1; i > 0; i = i - i&(-i) {
		sum += b[i]
		// fmt.Printf("index %v sum %v\n", i, sum)
	}
	// fmt.Println(".")
	return sum
}
