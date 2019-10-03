package main

import "fmt"

func main() {
	fmt.Println(pow(0, -4))
	fmt.Println(pow(2, -4))
	fmt.Println(pow(10, 4000000000))

}

func equal(a, b float64) bool {
	if a-b <= 0.000001 {
		return true
	}
	return false
}

func pow(value float64, step int) float64 {
	res := 1.0
	if equal(value, 0.0) {
		return 0
	}
	if step == 0 {
		if equal(value, 0.0) {
			return 0
		}
		return 1
	} else if step == 1 {
		return value
	}

	absStep := step
	if step < 0 {
		absStep = -step
	}
	tmpRes := res
	if absStep&1 == 1 {
		res *= value
		if res < tmpRes {
			return 0
		}
	}
	absStep >>= 1
	tmpRes = res
	res *= pow(value, absStep) * pow(value, absStep)
	if res < tmpRes {
		return 0
	}
	if step < 0 {
		return 1 / res
	}
	return res
}
