package str

import (
	"strconv"
)

func addBinary(a, b string, system int) string {
	result := ""
	flag := 0
	i,j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0{
		r1, r2 := 0, 0
		if i >= 0{
			r1 = int(a[i] - '0')
		}
		if j >= 0{
			r2 = int(b[j] - '0')
		}
		sum := r1 + r2 + flag
		mod := sum % system
		if mod != sum{
			flag = 1
		}else{
			flag = 0
		}
		result = strconv.Itoa(mod) + result
		i--
		j--
	}
	if flag != 0{
		result = strconv.Itoa(flag) + result
	}
	return result
}
