package piscine

import (
	"testing"
)

type testCase struct {
	name     string
	arg      []string
	expected []string
}

func TestKuromasu(t *testing.T) {
	testcases := []testCase{
		{
			name:     "Basic Test (PDF-1)",
			arg:      []string{"...2.", "..6.4", "5...6", "7.6..", ".3..."},
			expected: []string{".B.2B", "..6B4", "5B..6", "7.6B.", ".3B.."},
		},
		{
			name:     "Other Tests 1 (PDF-2)",
			arg:      []string{".....", "8.8..", ".7.7.", "..8.5", "....."},
			expected: []string{".B.B.", "8.8..", ".7.7B", "..8.5", "B.B.B"},
		},
		{
			name:     "Other Tests 2 (PDF-3)",
			arg:      []string{"37...", "..8..", ".....", "..8..", "...69"},
			expected: []string{"37.B.", "B.8..", "...B.", "B.8..", "...69"},
		},

		{
			name:     "Invalid Char",
			arg:      []string{"37.A.", "..8..", ".....", "..8..", "...69"},
			expected: []string{},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			result := Kuromasu(testcase.arg, false)

			if result == nil && len(testcase.expected) == 0 {
				return
			}
			if result == nil || len(result) != len(testcase.expected) {
				t.Errorf("NOT SAME LENGTH (expected height: %d, actual: %d)", len(testcase.expected), len(result))
			}

			for i, rowIntArr := range result {
				outstr := getRowStr(rowIntArr, "", true, "", false, SOLVER_WHITE_NUM, SOLVER_BLACK_NUM)

				if outstr != testcase.expected[i] {
					t.Errorf("Expected '%s' in [%d], buf got '%s' (GotString: |%s|).", testcase.expected, i, outstr, outstr)
				}
			}
		})
	}
}
