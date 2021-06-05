package lru_test

import (
	"github.com/scrocrix/algorithms/cache/lru"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LRUUnitTest struct {
	suite.Suite
}

func (unit *LRUUnitTest) TestCache() {

	// | Operations | 10  | 15  | 20  | 25  | 20 | 30  |
	// | Frame 1    | 10  | 10  | 20  | 20  | 20 | 20  |
	// | Frame 2    |     | 15  | 15  | 25  | 25 | 30  |
	// | Hit        | Yes | Yes | Yes | Yes | No | Yes |

	unit.Run("Replace item with least recently used", func() {
		sut := lru.NewLRU(2)

		sut.Cache(&lru.CacheItem{Key: "one", Value: "10"})
		sut.Cache(&lru.CacheItem{Key: "two", Value: "15"})
		sut.Cache(&lru.CacheItem{Key: "three", Value: "20"})
		sut.Cache(&lru.CacheItem{Key: "three", Value: "25"})
		sut.Cache(&lru.CacheItem{Key: "three", Value: "20"})
		sut.Cache(&lru.CacheItem{Key: "four", Value: "30"})

		require.Equal(unit.T(), "30", sut.GetItems().Front().Value.(*lru.CacheItem).Value)
		require.Equal(unit.T(), "20", sut.GetItems().Back().Value.(*lru.CacheItem).Value)
		require.Equal(unit.T(), 2, sut.GetItems().Len())
	})
}

func (unit *LRUUnitTest) TestGetItems() {
	unit.Run("Return items that have been cached", func() {
		sut := lru.NewLRU(1)

		require.Equal(unit.T(), 0, sut.GetItems().Len())

		sut.Cache(&lru.CacheItem{Key: "one", Value: "10"})

		require.Equal(unit.T(), 1, sut.GetItems().Len())

		require.Equal(unit.T(), "10", sut.GetItems().Front().Value.(*lru.CacheItem).Value)
		require.Equal(unit.T(), "one", sut.GetItems().Front().Value.(*lru.CacheItem).Key)
	})
}

func TestLRUUnitTest(t *testing.T) {
	suite.Run(t, new(LRUUnitTest))
}
