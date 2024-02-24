package atoi

const Zero = '0'
const Nine = '9'
const Eps = (1 << 31) - 1

func MyAtoi(s string) int {
	var result int
	var isDigit bool
	var sign int = 1

	for i := 0; i < len(s); i++ {
		b := s[i]
		// check sign
		if !isDigit && b == '-' {
			sign = -1
			isDigit = true
			continue
		}

		if !isDigit && b == '+' {
			sign = 1
			isDigit = true
			continue
		}

		if !isDigit && b >= Zero && b <= Nine {
			isDigit = true
		}

		if !isDigit && b != ' ' {
			return 0
		}

		if isDigit && (b < Zero || b > Nine) {
			return result * sign
		}

		if isDigit {
			// check upper limit
			if float64(Eps)/10.0-float64(result) < float64(int(b)-Zero)/10.0 && sign > 0 {
				return Eps
			}
			// check lower limit
			if float64(-Eps-1)/10.0+float64(result) > -float64(int(b)-Zero)/10.0 {
				return -Eps - 1
			}
			result = result*10 + (int(b) - Zero)
		}

	}
	return result * sign
}
