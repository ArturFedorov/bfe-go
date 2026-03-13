package find_median_from_data_stream

type MedianFinder struct{}

func Constructor() MedianFinder              { return MedianFinder{} }
func (mf *MedianFinder) AddNum(num int)      {}
func (mf *MedianFinder) FindMedian() float64 { return 0 }
