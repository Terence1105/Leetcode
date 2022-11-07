/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
  - Input: An array of strings strs

  - Output: Group the anagrams together.

  - 解法:
	step1: Declare map[string][]string{}(key儲存sorted string, value儲存unsort)
	step2: 依序將每個元素加入到對應的map string slice
	step3: 將 map 轉為 [][]string slice 並 return
    時間複雜度: O(N*(length of largest word))
    空間複雜度: O(n)
*/

// @lc code=start
func groupAnagrams(strs []string) [][]string {

	strsLen := len(strs)
	results := [][]string{}
	mapping := map[string][]string{}

	if strsLen == 0 {
		return results
	}

	for i := 0; i < strsLen; i++ {
		sortedStr := SortString(strs[i])
		mapping[sortedStr] = append(mapping[sortedStr], strs[i])
	}

	for _, value := range mapping {
		results = append(results, value)
	}

	return results
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// @lc code=end

