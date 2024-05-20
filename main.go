package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Input(r io.Reader) ([][]int, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, s.Err()
	}

	n, err := strconv.Atoi(s.Text())
	if err != nil {
		return nil, err
	}

	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		containers[i] = make([]int, 0, n)
		s.Scan()

		containerBalls := strings.Split(strings.Trim(s.Text(), " "), " ")
		if len(containerBalls) != n {
			return nil, fmt.Errorf("Incorrect number of balls in %d line", i+2)
		}

		for _, ball := range containerBalls {
			n, err := strconv.Atoi(ball)
			if err != nil {
				return nil, err
			}
			containers[i] = append(containers[i], n)
		}
	}
	return containers, nil

}

func Solve(containers [][]int) bool {
	n := len(containers)
	colors := make([]int, n)
	balls := make([]int, n)
	for i, container := range containers {
		for j, colorCount := range container {
			colors[j] += colorCount
			balls[i] += colorCount
		}

	}

	sort.Slice(colors, func(i, j int) bool {
		return colors[i] < colors[j]
	})
	sort.Slice(balls, func(i, j int) bool {
		return balls[i] < balls[j]
	})

	for i := 0; i < n; i++ {
		if colors[i] != balls[i] {
			return false
		}
	}
	return true
}

func main() {
	containers, err := Input(os.Stdin)
	if err != nil {
		panic(err)
	}

	if res := Solve(containers); res {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
