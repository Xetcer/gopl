package reverse

func reverse(pArr *[5]int) {
	for i, j := 0, len(pArr)-1; i < j; i, j = i+1, j-1 {
		pArr[i], pArr[j] = pArr[j], pArr[i]
	}
}
