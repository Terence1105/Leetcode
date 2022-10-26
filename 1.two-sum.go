/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 * Input: Nonsorted and duplicated one-dimension array
 * Output: Sorted one-dimension array contains two elements set
 * 解法:
   透過Dictionary解決
   Key儲存 target-nums[index]
   value儲存 index
   這樣只要 Dic[nums[i]]存在就找到答案了
   時間複雜度: O(n)
   空間複雜度: O(n)
*/

// @lc code=start
func twoSum(nums []int, target int) []int {

	numsCount := len(nums)
	mappingDic := map[int]int{}

	for i := 0; i < numsCount; i++ {
		_, isExist := mappingDic[nums[i]]
		if isExist {
			return []int{mappingDic[nums[i]], i}
		} else {
			mappingDic[target-nums[i]] = i
		}
	}
	return []int{0, 0}
}

// @lc code=end

