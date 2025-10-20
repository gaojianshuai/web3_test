// 考察：数字操作、条件判断
// 题目：判断一个整数是否是回文数

func isPalindrome(x int) bool {
	//转换成字符串
	str := strconv.Itoa(x)

	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true

}