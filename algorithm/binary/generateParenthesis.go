package binary

func generateParenthesis(n int) []string {
	var res []string
	/*
		判断回溯很简单，拿到一个问题，你感觉如果不穷举一下就没法知道答案，那就可以开始回溯了。

		一般回溯的问题有三种：

		1. Find a path to success 有没有解
		2. Find all paths to success 求所有解
		求所有解的个数
		求所有解的具体信息
		Find the best path to success 求最优解
		回溯法是一个剪枝了的二叉树。我们要得到的结果是可以 good leaf，如果不满足 good leaf 就继续向下搜索，搜索的时候需要满足一定的条件。
	*/

	var dfs func(x string, l, r int)
	dfs = func(x string, l, r int) {
		if l == 0 && r == 0 {
			res = append(res, x)
			return
		}

		if l > 0 {
			dfs(x+"(", l-1, r)
		}
		if l < r {
			dfs(x+")", l, r-1)
		}
	}
	// dfs ( "", 3,3)
	// dfs ( "(", 2,3)
	// dfs ( "((", 1,3)    // dfs ( "()", 2,2)

	// dfs ( "(((", 0,3)  // dfs ( "(()", 1,2) //dfs( "()(", 1,2)

	// dfs ( "((()", 0,2)
	// dfs ( "((())", 0,1)
	// dfs ( "((()))", 0,0)
	// -----1 -----

	dfs("", n, n)
	return res

}

// 电话号码2-9
func letterCombinations(digits string) []string {
	var r []string
	var dfs func(l []string, digits string) []string
	dfs = func(l []string, digits string) []string {
		// 终止条件
		// digist == ‘’
		if len(digits) == 0 {
			return l
		}

		// 单体
		l = cartesianProduct(l, phoneMap[string(digits[0])])

		return dfs(l, digits[1:])

	}
	return dfs(r, digits)

}

func cartesianProduct(a, b []string) []string {
	if a == nil {
		return b
	}
	var x []string
	for _, i := range a {
		for _, j := range b {
			x = append(x, i+j)
		}
	}
	return x
}

func helperX(s int32) []string {
	/*
		2 50
		9 57
	*/
	switch s {
	case 55:
		return []string{"p", "q", "r", "s"}
	case 56:
		return []string{"t", "u", "v"}
	case 57:
		return []string{"w", "x", "y", "z"}
	}
	var x []string
	for i := 0; i < 3; i++ {
		x = append(x, string((s-50)*3+int32(i)+97))
	}
	return x
}

var phoneMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "o", "n"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

var phoneMap2 map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap2[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
