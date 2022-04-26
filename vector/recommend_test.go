package vector

import (
	"math"
	"testing"
)

type UserLikeMusic struct {
	uid string

	// [music id]score
	// score higher, the more you like it
	musicLikeScore map[string]float64
}

type UserLikeMusicSet struct {
	// map[uid]index
	index map[string]int

	// contain all user info about music'like
	data []*UserLikeMusic
}

func NewUserLikeMusicSet() *UserLikeMusicSet {
	return &UserLikeMusicSet{
		index: make(map[string]int),
	}
}

func (u UserLikeMusicSet) Len() int {
	return len(u.index)
}

func (u UserLikeMusicSet) GetMatchDataByIndex(index int) MatchData {
	return u.data[index]
}

func (u *UserLikeMusicSet) Adds(ulmList ...*UserLikeMusic) {
	if u.index == nil {
		u.index = make(map[string]int)
	}
	for i, ulm := range ulmList {
		u.index[ulm.Id()] = i
		u.data = append(u.data, ulm)
	}
}

func (u UserLikeMusicSet) GetUserLikeMusices(uid string) []string {
	index := u.index[uid]
	md := u.GetMatchDataByIndex(index)

	musices := make([]string, len(md.Data()))
	for musicId := range md.Data() {
		musices = append(musices, musicId)
	}

	return musices
}

// O(n^2) = n(n-2)/2 ~= n^2
// return map[uid][recommend musicId...]
func RecommendMusicForOneUser(self UserLikeMusic, us UserLikeMusicSet) []string {
	uid := GetMostSimilar(self, us)
	return us.GetUserLikeMusices(uid)
	// TODO 求差集
}

// 所有用户都推荐
// 1. 设定每个人推荐多少首歌（10-50 首）
// 2. 每个人都推荐歌曲，1:[10,11,30,31] 排序（topK 堆排序）取最相似的几个人的歌单补全推荐Set

// 优化，提升时间 并发
// 优化，减少计算 3. 当一个人生成了推荐列表，其他和它相似的这几个人也应该有这个人的部分歌曲
// 优化，尽可能减少空间，共享

func (u UserLikeMusic) Id() string {
	return u.uid
}

func (u UserLikeMusic) Data() map[string]float64 {
	return u.musicLikeScore
}

var (
	self = UserLikeMusic{
		uid: "1",
		musicLikeScore: map[string]float64{
			"晴天":   5,
			"爱情买卖": -1,
			"轨迹":   2,
			"雅俗共赏": 5,
		},
	}

	others = NewUserLikeMusicSet()
)

func init() {
	others.Adds(
		&UserLikeMusic{uid: "1", musicLikeScore: map[string]float64{"晴天": 4, "爱情买卖": -1, "轨迹": 2, "雅俗共赏": 5}},
		&UserLikeMusic{uid: "2", musicLikeScore: map[string]float64{"晴天": 2, "爱情买卖": 5, "轨迹": 2, "雅俗共赏": -1}},
		&UserLikeMusic{uid: "2", musicLikeScore: map[string]float64{"晴天": 4, "爱情买卖": -1, "轨迹": 1, "雅俗共赏": 4}},
	)
}

func TestGetMostSimilar(t *testing.T) {
	type args struct {
		self   MatchData
		others MatchDataSet
	}
	tests := []struct {
		name   string
		args   args
		wantId string
	}{
		{name: "base", args: args{self: self, others: others}, wantId: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotId := GetMostSimilar(tt.args.self, tt.args.others); gotId != tt.wantId {
				t.Errorf("GetMostSimilar() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func TestCalDiff(t *testing.T) {
	type args struct {
		c1 MatchData
		c2 MatchData
	}
	tests := []struct {
		name     string
		args     args
		wantDiff float64
	}{
		{name: "base", args: args{c1: self, c2: others.GetMatchDataByIndex(0)}, wantDiff: math.Sqrt(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDiff := CalDiff(tt.args.c1, tt.args.c2); gotDiff != tt.wantDiff {
				t.Errorf("CalDiff() = %v, want %v", gotDiff, tt.wantDiff)
			}
		})
	}
}

// music test

// 测试优化，1.减少内存分配，共享 2.并发（加入分片锁）
