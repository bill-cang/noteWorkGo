package timeTools

import "fmt"

//判断year是否闰年
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%d年是闰年\n", year)
		return true
	} else {
		return false
	}
}


//求起止年份之间有多少个闰年
//startYear, endYear int 两个整型参数
func GetLeapYears(startYear, endYear int) (leapYears int) {
	for i := startYear; i < endYear+1; i++ {
		if IsLeapYear(i) {
			//leapYears=leapYears+1
			leapYears++
		}
	}
	return
}

