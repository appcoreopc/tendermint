package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlockchainInfo(t *testing.T) {

	cases := []struct {
		min, max       int64
		height         int64
		limit          int64
		newMin, newMax int64
		resultLength   int64
		wantErr        bool
	}{

		{1, 0, 5, 3, 2, 5, 3, false},
	} /*

		// min > max
		{0, 0, 0, 10, 0, true},  // min set to 1
		{0, 1, 0, 10, 0, true},  // max set to height (0)
		{0, 0, 1, 10, 1, false}, // max set to height (1)
		{2, 0, 1, 10, 0, true},  // max set to height (1)
		{2, 1, 5, 10, 0, true},

		// negative
		{-1, 10, 14, 10, 0, true},
		{10, -1, 14, 10, 0, true},
		{-9223372036854775808, -9223372036854775788, 100, 20, 0, true},

		// check limit and height
		{1, 5, 10, 10, 5, false},
		{1, 15, 10, 10, 10, false},
		{1, 15, 15, 10, 10, false},
		{1, 15, 15, 20, 15, false},

		{10, 0, 5, 0, 0, true},
		{10, 9, 5, 0, 0, true},

		//
	}*/

	for i, c := range cases {
		caseString := fmt.Sprintf("test %d failed", i)
		min, max, err := filterMinMax(c.height, c.min, c.max, c.limit)
		if c.wantErr {
			require.Error(t, err, caseString)
		} else {
			require.NoError(t, err, caseString)
			require.Equal(t, 1+max-min, c.resultLength, caseString)
		}
	}

}
