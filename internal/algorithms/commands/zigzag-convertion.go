package commands

type OnZigZagConvertor struct {
	Input   string
	NumRows int
}

func (zz *OnZigZagConvertor) Execute() string {
	if zz.NumRows == 1 || zz.NumRows >= len(zz.Input) {
		return zz.Input
	}

	rows := make([]string, zz.NumRows)
	currentRow := 0
	goingDown := false

	for _, char := range zz.Input {
		rows[currentRow] += string(char)
		if currentRow == 0 || currentRow == zz.NumRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			currentRow++
		} else {
			currentRow--
		}

	}
	result := ""

	for _, row := range rows {
		result += row
	}
	return result
}
