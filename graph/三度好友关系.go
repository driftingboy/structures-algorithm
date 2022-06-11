package graph

import (
	"container/list"
)

// 1.图的遍历
// 2. 二叉树的深度、广度优先遍和图的类似。
// 深度：借助一个栈（递归调用 或 自己实现栈效果）
// 广度：借助一个队列
type Graph struct {
	// 顶点的个数
	num int
	// 邻接表
	d map[string]*list.List
	// 反向邻接表
	reverse map[string]*list.List
}

func NewGraph() *Graph {
	return &Graph{
		d:       make(map[string]*list.List),
		reverse: make(map[string]*list.List),
	}
}

// 关注
func (g *Graph) Follow(id, followerId string) {
	l, ok := g.d[id]
	if !ok {
		l = list.New()
		g.d[id] = l
	}
	r, ok := g.reverse[followerId]
	if !ok {
		r = list.New()
		g.reverse[followerId] = r
	}

	g.num++
	l.PushBack(followerId)
	r.PushBack(id)
}

// 取消关注
func (g *Graph) UnFollow() {
}

// 查询关注的列表
func (g *Graph) ListFollowers(id string) []string {
	l, ok := g.d[id]
	if !ok {
		return []string{}
	}

	result := make([]string, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(string))
	}

	return result
}

// 查询粉丝
func (g *Graph) ListFans(id string) []string {
	r, ok := g.reverse[id]
	if !ok {
		return []string{}
	}

	result := make([]string, 0, r.Len())
	for e := r.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(string))
	}

	return result
}

// 查询是否是好友（关注与被关注）
func (g *Graph) IsFriend(id string, fid string) bool {
	return false
}

// 查询推荐关注（bfs n 层）
// 第 n 层
func (g *Graph) Recommend(uid string, n int) (ids []string) {
	if n < 1 {
		return []string{}
	}
	queue := make([]string, 1)
	queue[0] = uid
	visited := make(map[string]struct{})
	visited[uid] = struct{}{}
	// first layer
	layerNum := make(map[string]int, 0)
	layerNum[uid] = 0

	for len(queue) > 0 {
		id := queue[0]
		if layerNum[id] == n {
			return queue
		}
		queue = queue[1:]
		l, ok := g.d[id]
		if !ok {
			continue
		}
		for e := l.Front(); e != nil; e = e.Next() {
			eid := e.Value.(string)
			if _, ok := visited[eid]; ok {
				continue
			}
			// 记录id的层数，在大于层数n时退出
			layerNum[eid] = layerNum[id] + 1
			queue = append(queue, eid)
		}
	}

	return queue
}

var (
	found bool
)

// 查询关系链路（dfs）
func (g *Graph) Link(sourceId string, dstId string) []string {
	found = false
	visited := make(map[string]struct{}, g.num)
	pre := make([]string, 0, g.num)

	g.dfs(sourceId, dstId, visited, &pre)
	return pre
}

func (g *Graph) dfs(s, d string, visited map[string]struct{}, pre *[]string) {
	// 退出条件
	if found {
		return
	}
	visited[s] = struct{}{}
	if s == d {
		found = true
		return
	}

	if l, ok := g.d[s]; ok {
		for e := l.Front(); e != nil; e = e.Next() {
			v := e.Value.(string)
			if _, ok := visited[v]; ok || found {
				continue
			}
			*pre = append(*pre, v)
			g.dfs(v, d, visited, pre)
		}
	}
}
