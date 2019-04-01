package extdata

func StringArrayContains(sarr []string, v string) bool {
	for _, s := range sarr {
		if s == v {
			return true
		}
	}
	return false
}

func StringArrayReplace(sarr []string, c string, v string) bool {
	for i, s := range sarr {
		if s == c {
			sarr[i] = v
			return true
		}
	}
	return false
}
