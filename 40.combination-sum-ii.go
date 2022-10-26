/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {

	candidatesCount := len(candidates)
	answers := [][]int{}

	sort.Ints(candidates)

	for i := 0; i < candidatesCount; i++ {
		sum := candidates[i]
		answer := []int{sum}
		for j := i + 1; j < candidatesCount; j++ {
			if sum+candidates[j] == target {
				answer = append(answer, candidates[j])
				answers = append(answers, answer)
			} else if sum+candidates[j] < target {
				sum += candidates[j]
				answer = append(answer, candidates[j])
			}
		}
	}

	if candidates[candidatesCount-1] == target {
		answers = append(answers, []int{candidates[candidatesCount-1]})
	}

	return answers
}

func find
// @lc code=end

