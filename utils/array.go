package utils

// GenSep 生成sep字符的数组
func GenSep(num int, sep string) []string {
	results := make([]string, 0)
	for i := 0; i < num; i++ {
		results = append(results, sep)
	}

	return results
}
