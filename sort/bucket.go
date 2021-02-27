package sort

type Bucketsort interface {
	Sort(items []int) []int
}

type bucketsort struct {
	Bucketsort
	bucket        [][]int
	insertionSort Insertion
}

// NewBucketsort constructs a new instance of Bucketsort.
func NewBucketsort(insertionSort Insertion) Bucketsort {
	return &bucketsort{
		bucket:        make([][]int, 100),
		insertionSort: insertionSort,
	}
}

func (b *bucketsort) findBucketPositionForItem(item int) int {
	minRange := -10
	maxRange := -1
	bucketPosition := -1
	for item < minRange || item > maxRange {
		minRange = minRange + 10
		maxRange = maxRange + 10
		bucketPosition = bucketPosition + 1
		if item >= minRange && item <= maxRange {
			return bucketPosition
		}
	}
	return bucketPosition
}

func (b *bucketsort) Sort(items []int) []int {
	var sortedItems []int

	for _, item := range items {
		bucketPosition := b.findBucketPositionForItem(item)
		if bucketPosition != -1 {
			b.bucket[bucketPosition] = append(b.bucket[bucketPosition], item)
		}
	}

	for _, bucketItem := range b.bucket {
		if len(bucketItem) > 0 {
			si, _ := b.insertionSort.Sort(bucketItem, OrderAscending)
			sortedItems = append(sortedItems, si...)
		}
	}

	return sortedItems
}
