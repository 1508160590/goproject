package cal

func AddUpper(a int) int {
	res := 0
	for i := 0; i < a; i++ {
		res += i
	}
	return res
}
