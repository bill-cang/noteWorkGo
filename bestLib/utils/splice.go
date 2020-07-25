package utils

//求并集
func SlcUnion(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//求交集
func SlcIntersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func SlcDifference(sourceSlc, targetSlc []string) []string {
	m := make(map[string]int)
	df := make([]string, 0)
	sourceSlc = SlcDistinct(sourceSlc)
	inter := SlcIntersect(sourceSlc, targetSlc)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range sourceSlc {
		times, _ := m[value]
		if times == 0 {
			df = append(df, value)
		}
	}
	return df
}

//去重
func SlcDistinct(slc []string) (nSlc []string) {
	result := []string{}
	tempMap := map[string]byte{}
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l {
			result = append(result, e)
		}
	}
	return result
}

