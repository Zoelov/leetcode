package pascals_triangle

/******************************************************

https://leetcode-cn.com/problems/pascals-triangle/
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。


******************************************************/

func generate(numRows int) [][]int {
	ret := [][]int{}

	for j := 1; j <= numRows; j++ {

		data := []int{}
		for i := 0; i < j; i++ {
			if i == 0 || i == j-1 {
				data = append(data, 1)
			} else {
				v := ret[j-2][i] + ret[j-2][i-1]
				data = append(data, v)
			}

		}
		ret = append(ret, data)
	}

	return ret
}
