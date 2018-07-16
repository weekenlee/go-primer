package main

import (
	"sort"
	"fmt"
) 

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	record := make(map[string][]string)

	for _, str := range strs {
		tmp := sortString(str)
		record[tmp] = append(record[tmp], str)
	}

	for _, v :=range record {
		sort.Strings(v)
		res = append(res, v)
	}

	return res
}

func sortString(s string) string {
	bytes := []byte(s)

	temp := make([]int , len(bytes))
	for i, b := range bytes {
		temp[i] = int(b)
	}
	sort.Ints(temp)
	for i, v := range temp {
		bytes[i] = byte(v)
	}
	return string(bytes)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat","bat",}
	res := groupAnagrams(strs)
	for _, v :=range res {
		for _, vv := range v {
			fmt.Printf("%v ", vv)
		}
		fmt.Println("")
	}
}
