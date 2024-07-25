package commands

type OnLongestPalindrome struct {
	Input string
}

func (l *OnLongestPalindrome) Execute() string {
	if len(l.Input) == 0 {
		return ""
	}
	start, end := 0, 0

	for i := 0; i < len(l.Input); i++ {
		left1, right1 := l.expandFromCenter(l.Input, i, i)
		left2, right2 := l.expandFromCenter(l.Input, i, i+1)

		if right1-left1 > end-start {
			start, end = left1, right1
		}

		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}

	return l.Input[start : end+1]
}

func (l *OnLongestPalindrome) expandFromCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
