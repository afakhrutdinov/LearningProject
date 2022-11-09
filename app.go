package main

import (
	"fmt"
	"math"
)

type arr []int32

//type arr2 [][]int

type excludeType struct {
	i, j int
}

func main() {
	array := make(arr, 7)
	for i := 0; i < 7; i++ {
		array[0] = 1
		array[1] = 2
		array[2] = 1
		array[3] = 2
		array[4] = 1
		array[5] = 3
	}
	fmt.Println(reverseArray(array))
}

func hourglassSum(arr [][]int32) int32 {
	// Write your code here
	var maxCount int32

	return maxCount
}

func reverseArray(a []int32) []int32 {
	// Write your code here
	b := make([]int32, len(a))
	for i, _ := range a {
		b[i] = a[len(a)-i-1]
	}
	return b
}

type openedList struct {
	p1, p2 int32
}

func pageCount(n int32, p int32) int32 {
	// Write your code here
	array := make([]openedList, 0)
	page1 := int32(0)
	page2 := int32(0)
	pageCount := int32(0)
	pageGroup := 0
	for {
		page2 = page1 + 1
		pageCount += 2
		if pageCount > n+1 {
			page2 = 0
		}
		array = append(array, openedList{page1, page2})
		page1 += 2
		page2 += 2
		if pageCount >= n+1 {
			break
		}
	}
	for i, v := range array {
		if v.p1 == p || v.p2 == p {
			pageGroup = i
		}
	}
	fmt.Println(array)
	return int32(math.Min(float64(pageGroup), float64(len(array)-pageGroup-1)))
}

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	var pairsTotal int32
	pairsMap := make(map[int32]int32)
	for _, v := range ar {
		pairsMap[v]++
	}
	fmt.Println(pairsMap)
	for _, v := range pairsMap {
		pairsTotal += v / 2
	}
	return pairsTotal
}

type birds struct {
	btype, count int
}

func migratoryBirds(arr []int) int {
	maxValue := 0
	result := 1000000
	birds := make(map[int]int)
	for _, v := range arr {
		_, ok := birds[v]
		if ok == true {
			birds[v]++
		} else {
			birds[v] = 1
		}
	}
	for _, v := range birds {
		if v > maxValue {
			maxValue = v
		}
	}
	for i, v := range birds {
		if v != maxValue {
			delete(birds, i)
		}
	}
	for i := range birds {
		if i < result {
			result = i
		}
	}

	return result
}

func divisibleSumPairs(n int32, k int, ar []int) int {
	count := 0
	sum := 0
	for _, v1 := range ar {
		for _, v2 := range ar {
			if v1 < v2 {
				sum = v1 + v2
				if sum%k == 0 {
					//fmt.Printf("ar[%v] = %v, ar[%v] = %v, sum = %v, sum/k = %v\n", i, ar[i], j, ar[j], sum, sum%k)
					count++
				}
			}
		}
	}
	return count
}

func birthday1(s []int, d int, m int) int {
	result := 0
	sum := 0
	for i := 0; i < len(s); i++ {
		fmt.Println("___________", i)
		for j := 0; j < m; j++ {
			fmt.Println("___", i+j)
			if i+j < len(s) {
				sum += s[i+j]
			}
		}
		if sum == d {
			result++
		}
		sum = 0
	}
	return result
}

func birthday(s []int, d int, m int) int {
	s1 := s
	var count int
	exclude := make([]int, 0)
	var kapat bool

	for i, _ := range s {
		for j, _ := range s1 {
			if i != j && s[i]+s1[j] == d {
				for _, l := range exclude {
					if j == l {
						kapat = true
						break
					}
				}
				if kapat == false {
					count += 1
					exclude = append(exclude, i, j)
					fmt.Printf("s[%v]: %v, s1[%v]: %v, count: %v, exclude: %v\n", i, s[i], j, s1[j], count, exclude[1:])
				}
			}
		}
	}
	return count
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var countApples int32
	var countOranges int32

	for _, v := range apples {
		if v+a >= s && v+a <= t {
			countApples++
		}
	}
	for _, v := range oranges {
		if b-v <= t && b-v >= s {
			countOranges++
		}
	}
	fmt.Println(countApples)
	fmt.Println(countOranges)
}

func gradingStudents(grades []int) []int {
	// Write your code here
	var multiplesOfFive []int
	for i, j := 0, int(0); j <= 110; i++ {
		multiplesOfFive = append(multiplesOfFive, j)
		j += 5
	}
	var roundedValue int
	var flag bool = false
	roundedGrades := make([]int, len(grades))
	rVal := roundedValue
	for i, v := range grades {
		if v > 38 {
			for _, k := range multiplesOfFive {
				switch {
				case k == v:
					rVal = v
					flag = true
				case k > v:
					if k-v < 3 {
						rVal = k
						flag = true
					} else {
						rVal = v
						flag = true
					}
				}
				if flag == true {
					flag = false
					break
				}
			}
		} else {
			rVal = v
		}
		roundedGrades[i] = rVal
		rVal = 0
	}
	return roundedGrades
}

/*
	s := "12:05:45PM"
	ss := []byte(s)
	h, _ := strconv.Atoi(string(ss[:2]))
	var hh string
	if s[8:10] == "PM" && h < 12 {
		hh = strconv.Itoa(h + 12)
	} else if s[8:10] == "AM" && h == 12 {
		hh = "00"
	}
	if hh != "" {
		ss[0] = hh[0]
		ss[1] = hh[1]
	}
	fmt.Println(string(ss[:8]))
*/

/*
	s := "12:00:22PM"
	b := []byte(s)
	isPm := false
	for i, v := range b {
		if i == len(b)-2 {
			if strings.ToUpper(string(v)) == "P" {
				isPm = true
			}
		}
	}
	c := b[:len(b)-2]
	fmt.Println(string(c[0]) + string(c[1]))
	n := []byte(string(int(c[0])+1) + string(int(c[1])+2))
	if isPm == true {
		c[0] = n[0]
		c[1] = n[1]
		fmt.Println(string(c[0]) + string(c[1]))
		fmt.Println(int([]byte(string(c[0]) + string(c[1]))[0]))
		if l := int([]byte(string(c[0]) + string(c[1]))[0]); l == 24 {
			fmt.Println(l)
			c[0] = 0
			c[1] = 1
		}
	}
	fmt.Println(string(c))
	fmt.Printf("%T\n", string(c))

}
*/

func timeConversion(s string) string {
	// Write your code here
	//l := "15:04:05"
	//t, _ := time.Parse(l, s)
	return " "
	//switch {
	//case strings.Contains(s, "PM"):
	//	return "PM"
	//case strings.Contains(s, "AM"):
	//	return "AM"
	//}
	//return "Unknown format"
}

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	tNum := int32(0)
	tCount := int32(0)
	for _, v := range candles {
		if tNum <= v {
			tNum = v
		}
	}
	for _, v := range candles {
		if tNum == v {
			tCount += 1
		}
	}
	return tCount
}

func miniMaxSum(arr []int) {
	// Write your code here
	arrMin := 0
	arrMax := 0
	min := 0
	max := 0
	for i, v := range arr {
		fmt.Println(i, v)
		if i == 0 {
			min = v
		}
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	for _, v := range arr {
		if v != min {
			arrMax += v
		}
		if v != max {
			arrMin += v
		}
	}
	if arrMin == 0 {
		arrMin = min * (len(arr) - 1)
	}
	if arrMax == 0 {
		arrMax = max * (len(arr) - 1)
	}
	fmt.Println(arrMin, arrMax)
}

func staircase(n int32) {
	var str string
	for i := int32(0); i < n; i++ {
		for j := i; j < n+i; j++ {
			if j+1 < n {
				str = str + " "
			} else {
				str = str + "#"
			}
		}
		fmt.Println(str)
		str = ""
	}
}

func plusMinus(arr arr) {
	// Write your code here
	l := int32(len(arr))
	var p int32
	var n int32
	var z int32
	var infin float64
	var pos bool
	for _, v := range arr {
		infin = math.Inf(int(v))
		pos = math.IsInf(infin, 1)
		fmt.Println(v, infin, pos)
		switch {
		case v == 0:
			z += 1
		case pos == true:
			p += 1
		default:
			n += 1
		}
	}
	fmt.Println(fmt.Sprintf("%.4f", float64(p)/float64(l)))
	fmt.Println(fmt.Sprintf("%.4f", float64(n)/float64(l)))
	fmt.Println(fmt.Sprintf("%.4f", float64(z)/float64(l)))
}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	//var d1 int32
	//var d2 int32
	//for i, _ := range arr {
	//	fmt.Println(i)
	//	//for j, n := range arr[i] {
	//		//	d1 = n[i]
	//	}
	//
	//}
	return 0
}

//func diagonalDifference(arr [][]int) {
//	// Write your code here
//	//var d1 int32
//	//var d2 int32
//	for a, b := range arr {
//		fmt.Println(a, b)
//	}
