package problem0731

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
		start:  50,
		end:    60,
		result: true,
	},
	{
		start:  10,
		end:    40,
		result: true,
	},
	{
		start:  5,
		end:    15,
		result: false,
	},
	{
		start:  5,
		end:    10,
		result: true,
	},
	{
		start:  25,
		end:    55,
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
