package day_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_run_1(t *testing.T) {
	assert.EqualValues(t, []int{2, 0, 0, 0, 99}, run([]int{1, 0, 0, 0, 99}))
}

func Test_run_2(t *testing.T) {
	assert.EqualValues(t, []int{2, 3, 0, 6, 99}, run([]int{2, 3, 0, 3, 99}))
}

func Test_run_3(t *testing.T) {
	assert.EqualValues(t, []int{2, 4, 4, 5, 99, 9801}, run([]int{2, 4, 4, 5, 99, 0}))
}

func Test_run_4(t *testing.T) {
	assert.EqualValues(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, run([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}))
}
