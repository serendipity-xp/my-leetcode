package leetcode

/**
* @DateTime   : 2020/7/24 10:37
* @Author     : xulp
* @Description: 爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。

最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：

选出任一 x，满足 0 < x < N 且 N % x == 0 。
用 N - x 替换黑板上的数字 N 。
如果玩家无法执行这些操作，就会输掉游戏。

只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。
**/

func divisorGame(N int) bool {
	permit := 0
	for x := 1; x < N; x++ {
		if N%x == 0 {
			permit++
			N = N-x
			x = 0
		}
	}
	if permit%2 != 0 {
		return true
	}
	return false
}
// 0ms 1.9MB

func divisorGameByGuess(N int) bool {
	return N % 2 == 0
}