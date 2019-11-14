package main

import (
	"fmt"
	"strconv"
	"strings"
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

func spliceSum(s []uint16) (sum uint16) {
	for _, i := range s {
		if i == 0 {
			return sum
		}
		sum += i
	}
	return sum
}

func spliceSearch(n uint16, s []uint16) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}
	return false
}

/*func ogrv2Internal(length, level int, max *int, l, d []int) {
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
*/

var countAll, countSum, prev uint64
var dist uint16
var max uint16

func ogrv3Internal(length, level, sum, f uint16, l, d []uint16) {
	if level == length-1 {
		countAll++
		if countSum/1048576/4 > prev {
			prev = countSum / 1048576 / 4
			fmt.Printf("  all: %d, sum: %d, time: %v, current: %v\n", countAll, countSum, time.Now(), l)
		}
		if sum <= max {
			max = sum
			fmt.Printf("%d: %v\n", sum, l)
		}
		return
	}
	if sum > max {
		countSum++
		return
	}

	lo := level * (level + 1) / 2
	hi := lo + level
ext:
	for n := uint16(1); n < length*f; n++ {
		if !spliceSearch(n, d[:lo]) {
			l[level] = n
			dist = n
			d[lo] = n
			for i := lo + 1; i <= hi; i++ {
				dist += l[hi-i]
				if spliceSearch(dist, d[:lo]) {
					continue ext
				}
				d[i] = dist
			}
			ogrv3Internal(length, level+1, sum+n, 3, l, d)
		}
	}
}

func ogrv3(length int) {
	d := make([]uint16, length*(length-1)/2)
	l := make([]uint16, length-1)
	max = uint16(length * length)
	ogrv3Internal(uint16(length), 0, 0, 1, l, d)
}

func maxdist(s string) int {
	max := 0
	prev := 0
	for _, item := range strings.Split(s, " ") {
		n, err := strconv.ParseInt(item, 10, 32)
		if err == nil {
			if max < int(n)-prev {
				max = int(n) - prev
				//fmt.Printf("%d-%d: %d\n", prev, n, max)
			}
			prev = int(n)
		}
	}
	return max
}

func main() {
	//ogrv1()
	startTime := time.Now()
	ogrv3(10)
	endTime := time.Now()
	fmt.Printf("start: %v\nend  : %v\n", startTime, endTime)
	//fmt.Println(maxdist("0 1 8 11 68 77 94 116 121 156 158 179 194 208 212 228 240 253 259 283"))
	//fmt.Println(maxdist("0 2 24 56 77 82 83 95 129 144 179 186 195 255 265 285 293 296 310 329 333"))
	//fmt.Println(maxdist("0 1 9 14 43 70 106 122 124 128 159 179 204 223 253 263 270 291 330 341 353 356"))
	//fmt.Println(maxdist("0 3 7 17 61 66 91 99 114 159 171 199 200 226 235 246 277 316 329 348 350 366 372"))
	//fmt.Println(maxdist("0 9 33 37 38 97 122 129 140 142 152 191 205 208 252 278 286 326 332 353 368 384 403 425"))
	//fmt.Println(maxdist("0 12 29 39 72 91 146 157 160 161 166 191 207 214 258 290 316 354 372 394 396 431 459 467 480"))
	//fmt.Println(maxdist("0 1 33 83 104 110 124 163 185 200 203 249 251 258 314 318 343 356 386 430 440 456 464 475 487 492"))
	//fmt.Println(maxdist("0 3 15 41 66 95 97 106 142 152 220 221 225 242 295 330 338 354 382 388 402 415 486 504 523 546 553"))
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
