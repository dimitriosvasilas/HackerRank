package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var productOfRestTests = []struct {
	array            []int
	expectedResponse []int
}{
	{
		[]int{},
		[]int{},
	},
	{
		[]int{1, 2, 3, 4, 5},
		[]int{120, 60, 40, 30, 24},
	},
	{
		[]int{3, 2, 1},
		[]int{2, 3, 6},
	},
}

func TestProductOfRest(t *testing.T) {
	for _, tt := range productOfRestTests {
		response := productOfRest(tt.array)
		assert.Equal(t, tt.expectedResponse, response, "")
	}
}
