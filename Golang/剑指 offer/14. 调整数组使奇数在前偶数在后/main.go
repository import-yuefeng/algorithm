package main

func main() {

}

func oddEven(num []int) {
	l, r := 0, len(num)-1
	for l < r {
		for l < r && num[l]&1 == 1 {
			l++
		}
		for l < r && num[r]&1 == 0 {
			r--
		}
		if l >= r {
			break
		}
		num[l] += num[r]
		num[r] = num[l] - num[r]
		num[l] -= num[r]
	}
}
