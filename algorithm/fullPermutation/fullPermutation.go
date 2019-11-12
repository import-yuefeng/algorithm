// package fullPermutation
package main

import "fmt"

func main() {
	str := []string{"A", "B", "C", "D"}
	// res := outOrder(str)
	res := insert(str[:len(str)-1], str[len(str)-1])
	fmt.Println(res)

}

func insert(res []string, insertNum string) []string {
	result := make([]string, len(res)*(len(res[0])+1))
	index := 0
	for _, v := range res {
		fmt.Println(res)
		for i := 0; i < len(v); i++ {
			result[index] = v[:i] + insertNum + v[i:]
			index++
		}
		result[index] = v + insertNum
		index++
	}
	return result
}

func outOrder(str []string) []string {
	count := len(str)
	if count == 0 {
		return []string{}
	}
	if count == 1 {
		return str
	}
	return insert(outOrder(str[:count-1]), str[count-1])
}
