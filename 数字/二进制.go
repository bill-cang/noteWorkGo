package nums

import "strconv"

func ToBinary(ni int) string{
	var res  string
	var yu int =0
	for i:=0;i<10000 ;i++  {
		yu = ni/2
		mo := ni%2

		if yu == 0 {
			res =  strconv.Itoa(mo) + res
			return res
		}else {
			//除2求余数时，当模为1时，要减模
			switch mo {
			case 0:
				ni = ni -yu
			case 1:
				ni = ni -yu -1
			}

			if i != 0 && i%4 == 0{
				res =  strconv.Itoa(mo)+ "," + res
			}else{
				res =  strconv.Itoa(mo) + res
			}
		}
	}
	return "异常"
}

