package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsPairSum = []struct {
	array            []int
	k                int
	expectedResponse bool
}{
	{
		[]int{},
		42,
		false,
	},
	{
		[]int{10, 15, 3, 7},
		17,
		true,
	},
	{
		[]int{1, 4, 45, 6, 10, -8},
		16,
		true,
	},
	{
		[]int{1, 4, 45, 6, 10, -8},
		1,
		false,
	},
}

func TestPairSum(t *testing.T) {
	for _, tt := range testsPairSum {
		response := pairSumHash(tt.array, tt.k)
		assert.Equal(t, tt.expectedResponse, response, "")
		response = pairSumSort(tt.array, tt.k)
		assert.Equal(t, tt.expectedResponse, response, "")
	}
}
