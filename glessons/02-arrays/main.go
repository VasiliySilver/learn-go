package main

func main() {
	var scores [5]int
	scores[0] = 98
	scores[1] = 76
	scores[2] = 54
	scores[3] = 23
	scores[4] = 45

	newScores := []int{98, 76, 54, 23, 45}

	newScores = append(newScores, 100, 200, 300)

	println("Scores array length:", len(scores))
	println("NewScores slice length:", len(newScores))

	for i, score := range newScores {
		println("Score:", score, "index:", i)
	}
}
