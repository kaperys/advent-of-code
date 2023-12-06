package main

func main() {
	total := 1

	times := []int{53, 71, 78, 80}
	distances := []int{275, 1181, 1215, 1524}

	for race := 0; race < len(times); race++ {
		var count int

		for speed := 1; speed < times[race]; speed++ {
			if (times[race]-speed)*speed > distances[race] {
				count++
			}
		}

		total *= count
	}

	println(total)
}
