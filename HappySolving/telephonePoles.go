package main
import ("fmt")

var A, W []int
var cache [][]int


func main() {
	var N,K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	A = make([]int, N)
	W = make([]int, N)
	for n := range A {
		var x,y int
		fmt.Scanf("%d %d", &x, &y)
		A[n] = x
		W[n] = y
	}
	cache = make([][]int, N + 1)
	for n := 1; n <= N; n++ {
		cache[n] = make([]int, K + 1)
	}
	fmt.Println(cost(N, K))

}

func cost(n, k int) int {
	if n == k {
		return 0
	}
	if k <= 0 {
		return 1e15
	}
	if  cache[n][k]> 0 {
		return cache[n][k]
	}
	best := int(1e14)
	spent := 0
	weight := 0
	for u := n - 1;; u-- {
		c := spent + cost(u, k - 1)
		if c < best {
			best = c
		}
		if u == k - 1 {
			break
		}
		weight += W[u]
		spent += weight * (A[u] - A[u-1])
		if spent >= best {
			break
		}
	}
	cache[n][k] = best
	return best
}



