package main

func main() {
	var total int

	time := 53717880
	distance := 275118112151524

	for speed := 1; speed < time; speed++ {
		if (time-speed)*speed > distance {
			total++
		}
	}

	println(total)
}
