package main

import (
	"fmt"
	"time"
)

func ogrv1() {
	ogrLen := 4
	bestLen := uint16(999)
	buf := make([]uint16, ogrLen)
	lens := make([]uint16, ogrLen*(ogrLen-1)/2)
	//var best []uint16
	var baseindex int
	//var currentLen uint16
	for {
		for i := 1; i < ogrLen; i++ {
			baseindex = i * (i - 1) / 2
			if i != 1 {
				buf[i] = buf[i-1]
			}
		check1:
			for {
				buf[i]++
				for k := 0; k < i; k++ {
					lens[baseindex+k] = buf[i] - buf[k]
				}
				for k := 0; k < baseindex+i-1; k++ {
					for n := k + 1; n < baseindex+i; n++ {
						if lens[k] == lens[n] {
							continue check1
						}
					}
				}
				break
			}
		}
		if buf[ogrLen-1] <= bestLen {
			bestLen = buf[ogrLen-1]
			fmt.Printf("%v\n", buf)
		}
		if int(buf[ogrLen-1]) > ogrLen*ogrLen*ogrLen {
			break
		}
	}
}

func spliceSum(s []int) (sum int) {
	for _, i := range s {
		if i == 0 {
			return sum
		}
		sum += i
	}
	return sum
}

func spliceSearch(n int, s []int) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}
	return false
}

func ogrv2Internal(length, level int, max *int, l, d []int) {
	if level == length-1 {
		if curLen := spliceSum(l); curLen < *max {
			*max = curLen
			fmt.Printf("%d: %v\n", curLen, l)
		}
		return
	}
	if spliceSum(l[:level]) > *max {
		return
	}

	lo := level * (level + 1) / 2
	hi := lo + level + 1
ext:
	for n := 1; n < length*2; n++ {
		if !spliceSearch(n, d[:lo]) {
			l[level] = n
			for i := lo; i < hi; i++ {
				dist := spliceSum(l[i-lo : level+1])
				if spliceSearch(dist, d[:lo]) {
					continue ext
				}
				d[i] = dist
			}
			ogrv2Internal(length, level+1, max, l, d)
		}
	}
}

func ogrv2(length int) {
	d := make([]int, length*(length-1)/2)
	l := make([]int, length-1)
	max := length * length
	ogrv2Internal(length, 0, &max, l, d)
}

func main() {
	//ogrv1()
	startTime := time.Now()
	ogrv2(11)
	endTime := time.Now()
	fmt.Printf("start: %v\nend  : %v\n", startTime, endTime)
}

/*
d1
d1+d2
d1+d2+d3
d1+d2+d3+d4
d2
d2+d3
d2+d3+d4
d3
d3+d4
d4
*/
