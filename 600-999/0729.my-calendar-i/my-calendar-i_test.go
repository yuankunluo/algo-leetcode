package problem0729

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCase = []struct {
	start, end int
	result     bool
}{
	{
		start:  10,
		end:    20,
		result: true,
	},
	{
		start:  15,
		end:    25,
		result: false,
	},
	{
		start:  20,
		end:    30,
		result: true,
	},
}

func TestBook(t *testing.T) {
	ast := assert.New(t)

	obj := Constructor()

	for _, tc := range testCase {
		ast.Equal(tc.result, obj.Book(tc.start, tc.end))
	}
}
