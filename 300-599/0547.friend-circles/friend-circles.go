

func findCircleNum(M [][]int) int {

	n := len(M)
	if n == 0 {
		return 0
	}

	visited := make([]int, n)
	ans := 0
	

	for i := 0 ; i < n ; i++ {
		if visited[i] == 1 {
			continue
		}
		dfs(M, i, n, visited)
		ans++
	}
	return ans
}


func dfs(M [][]int, cur int, n int, visited []int){
	if visited[cur] == 1{
		return
	}
	visited[cur] = 1
	// Visit all friends 
	for i := 0; i < n ; i++ {
		if M[cur][i] == 1 && visited[i] == 0{
			dfs(M, i, n, visited)
		}
	}
}