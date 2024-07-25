package commands

type OnLongestSubstring struct {
	Input string
}

func (ls *OnLongestSubstring) Execute() int {
	lastSeen := make(map[string]int)
	start, maxLength, maxStart := 0, 0, 0

	for i := 0; i < len(ls.Input); i++ {
		if lastPos, found := lastSeen[string(ls.Input[i])]; found && lastPos >= start {
			start = lastPos + 1
		}
		lastSeen[string(ls.Input[i])] = i
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
			maxStart = start
		}
	}
	return len(ls.Input[maxStart : maxStart+maxLength])
}
