package main

import "sort"

const VALUE byte = '0'

func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	set := make(map[int]byte)
	var result []int
	for _,v := range ns {
		if _, ok := set[v]; !ok {
			set[v] = VALUE
			result = append(result, v)
		}
	}
	sort.Ints(result)
	n := len(result)
	for i := 0; i < n; i++ {
		cross(result[i], len(xs), xs, ys)
	}
	return xs, ys
}

func cross(start int, end int, xs []int, ys []int) {
	for j := start; j < end; j++ {
		xs[j], ys[j] = ys[j], xs[j]
	}
}
