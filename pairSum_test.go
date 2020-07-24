package algo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	returnCode := m.Run()
	os.Exit(returnCode)
}

var tests = []struct {
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
	for _, tt := range tests {
		response := pairSumHash(tt.array, tt.k)
		assert.Equal(t, tt.expectedResponse, response, "")
		response = pairSumSort(tt.array, tt.k)
		assert.Equal(t, tt.expectedResponse, response, "")
	}
}
