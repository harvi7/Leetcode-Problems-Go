func countComponents(n int, edges [][]int) int {
    dsu := make([]int, n)
    for i := range dsu {
        dsu[i] = -1
    }    
    
    for _, edge := range edges {
        if union(dsu, edge[0], edge[1]) {
            n--
        }
    }
    return n
}

func find(dsu []int, x int) int {
    if dsu[x] < 0 {
        return x
    }
    
    dsu[x] = find(dsu, dsu[x])
    
    return dsu[x]
}

func union(dsu []int, u, v int) bool {
    pu, pv := find(dsu, u), find(dsu, v)
    
    if pu == pv {
        return false
    }
    
    if dsu[pu] < dsu[pv] {
        dsu[pu] += dsu[pv]
        dsu[pv] = pu
    } else {
        dsu[pv] += dsu[pu]
        dsu[pu] = pv
    }
    
    return true
}
