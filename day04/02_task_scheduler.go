package main

import (
	"fmt"
	"sort"
)

func task_scheduler(tasks []rune, distance int) int {
	freq_map := map[rune]int {}

	for _, c := range tasks {
		freq_map[c] += 1
	}

	frequencies := []int {}

	for k, _ := range freq_map {
		frequencies = append(frequencies, freq_map[k])
	}

	sort.Ints(frequencies)

	max_freq := frequencies[len(frequencies) - 1]

	idle_time := (max_freq - 1) * distance

	fmt.Println(frequencies, max_freq, idle_time)

	for _, f := range frequencies[ : len(frequencies) - 1 ] {
		idle_time -= f
	}

	if idle_time < 0 {
		idle_time = 0
	}

	return idle_time + len(tasks)
}

func main() {

	test_data := []rune {
		'A', 'B', 'A', 'B', 'C', 'A',
	}

	fmt.Println(task_scheduler(test_data, 2))
}
