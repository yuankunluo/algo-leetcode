package problem0190

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  uint32
	output uint32
}{
	{
		0b00000010100101000001111010011100,
		0b00111001011110000010100101000000,
	},
	{
		0b11111111111111111111111111111101,
		0b10111111111111111111111111111111,
	},
}

func TestReverseBits(t *testing.T) {
	ast := assert.New(t)
	for _, ts := range testCases {
		ast.Equal(reverseBits(ts.input), ts.output)
	}
}
