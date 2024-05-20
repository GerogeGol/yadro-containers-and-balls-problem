package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		input [][]int
		want  bool
	}{
		{
			input: [][]int{
				{1, 2},
				{2, 1}},
			want: true,
		},
		{
			input: [][]int{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 1},
			},
			want: false,
		},
		{
			input: [][]int{
				{1, 0},
				{1, 1},
			},
			want: true,
		},
		{
			input: [][]int{
				{1, 0},
				{1, 2},
			},
			want: false,
		},
		{
			input: [][]int{
				{5, 10, 5},
				{10, 5, 5},
				{5, 5, 0},
			},
			want: true,
		},
		{
			input: [][]int{
				{20, 30},
				{1000000000, 20},
			},
			want: true,
		},
		{
			input: [][]int{
				{20, 30},
				{1000000000, 21},
			},
			want: false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			got := Solve(c.input)
			if got != c.want {
				t.Errorf("got %t, want %t", got, c.want)
			}
		})
	}
}

func TestInput(t *testing.T) {
	cases := []struct {
		input string
		want  [][]int
	}{
		{
			input: `2
1 2
2 1
	`,
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			input: `3
10 20 30
1 1 1
0 0 1
	`,
			want: [][]int{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 1},
			},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			buf := bytes.NewBuffer([]byte(c.input))
			got, err := Input(buf)
			assertNoError(t, err)

			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %#v, want %#v", got, c.want)
			}
		})
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
}
