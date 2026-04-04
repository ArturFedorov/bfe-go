package redundant_connection

type UnionFind struct {
	parents []int
}

func newUnionFind(size int) *UnionFind {
	u := &UnionFind{parents: make([]int, size+1)} //// size+1 so index 1..n valid

	for i := range u.parents {
		u.parents[i] = i
	}

	return u
}

func (u *UnionFind) find(i int) int {
	for u.parents[i] != i {
		u.parents[i] = u.parents[u.parents[i]] // path compression
		i = u.parents[i]
	}
	return i
}

func (u *UnionFind) union(a, b int) bool {
	pA, pB := u.find(a), u.find(b)

	if pA == pB {
		return false
	}
	if pA < pB {
		u.parents[pB] = pA
	} else {
		u.parents[pA] = pB
	}

	return true
}

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	uf := newUnionFind(n)

	for _, edge := range edges {
		if !uf.union(edge[0], edge[1]) {
			return edge
		}
	}

	return nil
}
