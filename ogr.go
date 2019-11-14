package ogr

// <editor-fold defaultstate="collapsed">
/* func Ogrv1(ogrLen int) {
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
*/ // </editor-fold>

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
	countSplice++
	for _, i := range s {
		if i == n {
			return true
		}
	}
	return false
}

func ogrv2Internal(length, level uint16, max *uint16, l, d []uint16) {
	if level == length-1 {
		if curLen := spliceSum(l); curLen < *max {
			*max = curLen
			//fmt.Printf("%d: %v\n", curLen, l)
			best = append([]uint16(nil), l...)
		}
		return
	}
	if spliceSum(l[:level]) > *max {
		return
	}

	lo := level * (level + 1) / 2
	hi := lo + level + 1
ext:
	for n := uint16(1); n < length*2; n++ {
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

func Ogrv2(length int) []uint16 {
	d := make([]uint16, length*(length-1)/2)
	l := make([]uint16, length-1)
	max := uint16(length * length)
	ogrv2Internal(uint16(length), 0, &max, l, d)
	return best
}

var countAll, countSum, countSplice, prev uint64
var dist uint16
var max uint16
var best []uint16

func ogrv3Internal(length, level, sum, f uint16, l, d []uint16) {
	if level == length-1 {
		countAll++
		if countSum/1048576/4 > prev {
			prev = countSum / 1048576 / 4
			//fmt.Printf("  all: %d, sum: %d, splice: %d, time: %v, current: %v\n", countAll, countSum, countSplice, time.Now(), l)
		}
		if sum < max {
			max = sum
			//fmt.Printf("%d: %v\n", sum, l)
			best = append([]uint16(nil), l...)
		}
		return
	}
	if sum > max {
		countSum++
		return
	}

	lo := level * (level + 1) / 2
	hi := lo + level
	spliceLookup := make([]bool, 1024)
	for _, item := range d[:lo] {
		spliceLookup[item] = true
	}
ext:
	for n := uint16(1); n < length*f; n++ {
		if !spliceLookup[n] {
			l[level] = n
			dist = n
			d[lo] = n
			for i := lo + 1; i <= hi; i++ {
				dist += l[hi-i]
				if spliceLookup[dist] {
					continue ext
				}
				d[i] = dist
			}
			ogrv3Internal(length, level+1, sum+n, 3, l, d)
		}
	}
}

func Ogrv3(length int) []uint16 {
	d := make([]uint16, length*(length-1)/2)
	l := make([]uint16, length-1)
	max = uint16(length * length)
	ogrv3Internal(uint16(length), 0, 0, 1, l, d)
	return best
}
