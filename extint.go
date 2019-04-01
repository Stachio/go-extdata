package extdata

//Courtesy of https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go#answer-6878625
//@nmichaels
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

//--------------------------------------------------------------------------------------------------------------

func IntArrayContains(iarr []int, v int) bool {
	for _, n := range iarr {
		if n == v {
			return true
		}
	}
	return false
}

//Returns the index and value of the greatest value
func IntArrayGreatest(iarr []int) (i int, v int) {
	i = 0
	v = MinInt
	for j, n := range iarr {
		if n > v {
			i = j
			v = n
		}
	}
	return
}
