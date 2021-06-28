package phonenumber_test

import (
	"github.com/scrocrix/algorithms/phonenumber"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindMatches(t *testing.T) {
	t.Run("Return possible strings", func(t *testing.T) {
		matches := phonenumber.FindMatches("3662277", []string{
			"foo",
			"bar",
			"baz",
			"foobar",
			"emo",
			"cap",
			"car",
			"cat",
		})

		require.Equal(t, []string{"bar", "cap", "car", "emo", "foo", "foobar"}, matches)
	})
}
