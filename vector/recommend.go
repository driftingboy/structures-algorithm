package vector

import "math"

type MatchData interface {
	Id() string
	Data() map[string]float64 // 或者使用一个map数据可以无序
}

type MatchDataSet interface {
	Len() int
	GetMatchDataByIndex(index int) MatchData
}

func CalDiff(c1, c2 MatchData) (diff float64) {
	matchkeys := keySet(c1, c2)

	var sum float64
	for _, k := range matchkeys {
		sum += math.Pow(c1.Data()[k]-c2.Data()[k], 2)
	}

	return math.Sqrt(sum)
}

func keySet(c1, c2 MatchData) []string {
	maxLength := len(c1.Data())
	if len(c2.Data()) > maxLength {
		maxLength = len(c2.Data())
	}

	keySet := make(map[string]struct{}, maxLength)
	for k := range c1.Data() {
		keySet[k] = struct{}{}
	}
	for k := range c2.Data() {
		keySet[k] = struct{}{}
	}

	keySetResult := make([]string, maxLength)
	for k := range keySet {
		keySetResult = append(keySetResult, k)
	}

	return keySetResult
}

func GetMostSimilar(self MatchData, others MatchDataSet) (id string) {
	var minDiff = math.MaxFloat64

	for i := 0; i < others.Len(); i++ {
		o := others.GetMatchDataByIndex(i)
		curDiff := CalDiff(self, o)
		if curDiff < minDiff {
			minDiff = curDiff
			id = o.Id()
		}
	}

	return
}
