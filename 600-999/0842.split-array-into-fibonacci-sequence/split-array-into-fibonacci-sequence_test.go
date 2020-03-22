package problem0842

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	output []int
}{
	{
		input: "123456579",
		output: []int{
			123,
			456,
			579,
		},
	},
	{
		input:  "502113822114324228146342470570616913086148370223967883880490627727810157768164350462659281443027860696206741126485341822692082949177424771869507721046921249291642202139633432706879765292084310",
		output: []int{},
	},
}

func TestSplitIntoFibonacci(t *testing.T) {
	ast := assert.New(t)

	for _, tsc := range testCases {
		ast.Equal(tsc.output, splitIntoFibonacci(tsc.input))
	}
}
