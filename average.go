package main

import "fmt"

func main() {
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

	// calc average array
	// lengthArray := len(score)
	totalScore := 0
	for i := 0; i < len(scores); i++ {
		totalScore = totalScore + scores[i]
	}

	averageScore := float64(totalScore) / float64(len(scores))
	fmt.Println("sum score : ", totalScore)
	fmt.Println("average score : ", float64(averageScore))

}
