package day01

import (
	"fmt"
	"strconv"
)

func Run(lines []string) error {
	values := make([]int, 0)
	for _, line := range lines {
		converted, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("error converting number '%s': %v", line, err)
		}
		values = append(values, converted)
	}

	x, y, err := findPart1(values)
	if err != nil {
		return err
	}
	fmt.Println("Part 1:", x*y)

	x, y, z, err := findPart2(values)
	if err != nil {
		return err
	}
	fmt.Println("Part 2:", x*y*z)

	return nil
}

func findPart1(values []int) (int, int, error) {
	for _, x := range values {
		for _, y := range values {
			if x+y == 2020 {
				return x, y, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("no values summed to 2020")
}

func findPart2(values []int) (int, int, int, error) {
	for _, x := range values {
		for _, y := range values {
			for _, z := range values {
				if x+y+z == 2020 {
					return x, y, z, nil
				}
			}
		}
	}
	return 0, 0, 0, fmt.Errorf("no values summed to 2020")
}
