package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

type sourceDestMap struct {
	sourceStart, sourceEnd int64
	destStart, destEnd     int64
	sRange                 int64
}

func main() {
	input, err := os.ReadFile("2023/05/input.txt")
	if err != nil {
		panic(err)
	}

	var (
		seeds []int64
		maps  map[string][]sourceDestMap = make(map[string][]sourceDestMap)
	)

	for _, part := range strings.Split(string(input), "\n\n") {
		switch {
		case len(part) > 4 && part[0:5] == "seeds":
			for _, seed := range strings.Fields(part[7:]) {
				sn, _ := strconv.ParseInt(seed, 10, 64)
				seeds = append(seeds, sn)
			}
		default:
			lines := strings.Split(part, "\n")
			name := lines[0][:len(lines[0])-5]

			for _, line := range lines[1:] {
				if len(line) > 0 {
					nums := strings.Fields(line)

					destStart, _ := strconv.ParseInt(nums[0], 10, 64)
					sourceStart, _ := strconv.ParseInt(nums[1], 10, 64)
					sRange, _ := strconv.ParseInt(nums[2], 10, 64)

					maps[name] = append(maps[name], sourceDestMap{
						destStart:   destStart,
						destEnd:     destStart + (sRange - 1),
						sourceStart: sourceStart,
						sourceEnd:   sourceStart + (sRange - 1),
						sRange:      sRange,
					})
				}
			}
		}
	}

	var (
		destNum   int64
		sourceNum int64

		locations []int64
	)

	for _, seed := range seeds {
		sourceNum = seed

		for _, category := range []string{
			"seed-to-soil",
			"soil-to-fertilizer",
			"fertilizer-to-water",
			"water-to-light",
			"light-to-temperature",
			"temperature-to-humidity",
			"humidity-to-location",
		} {
			for _, m := range maps[category] {
				destNum = sourceNum

				if sourceNum >= m.sourceStart && sourceNum <= m.sourceEnd {
					dest := m.destStart + (sourceNum - m.sourceStart)

					if dest >= m.destStart && dest <= m.destEnd {
						destNum = dest
						sourceNum = destNum

						break
					}
				}

				sourceNum = destNum
			}
		}

		locations = append(locations, destNum)
	}

	println(slices.Min(locations))
}
