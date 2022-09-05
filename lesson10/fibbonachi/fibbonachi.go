package fibbonachi

func fibbonachi(n uint64) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibbonachi(n-1) + fibbonachi(n-2)
	}
}
