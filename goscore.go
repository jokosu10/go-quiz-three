package main

import "fmt"

func unmain() {
	scores := [8]int{
		100,
		80,
		75,
		92,
		70,
		93,
		88,
		67,
	}

	var goodScore []int

	for _, value := range scores {
		if value >= 90 {
			goodScore = append(goodScore, value)
		}
	}

	for _, value := range goodScore {
		fmt.Println(value)
	}
	fmt.Println(goodScore)
}
