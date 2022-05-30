package piscine

// []stringを操作する関数
// 各要素のlenは等しいことを仮定する

func TrimBlankRowCol(lines []string) []string {
	return TransposeLines(TrimBlankLine(TransposeLines(TrimBlankLine(lines))))
}

func TrimBlankLine(lines []string) []string {
	f := 0
	t := len(lines) - 1
	for ; f <= t; f++ {
		if StringHasRune(lines[f], '#') {
			// lines[f]が空行でない
			break
		}
	}
	for ; f <= t; t-- {
		if StringHasRune(lines[t], '#') {
			// lines[t]が空行でない
			break
		}
	}
	// ->
	// lines[f], lines[t]はともに空行ではない
	// lines[f-1], lines[t+1]は存在するなら空行
	return lines[f : t+1]
}

// 左右反転
// reverse every lines
func ReverseLines(lines []string) []string {
	return Map(lines, func(line string, _ int) string {
		return string(Reverse([]rune(line)))
	})
}

// 回転
// 時計回りに90度回す
// rotate 90degrees in clockwise
func RotateLines(lines []string) []string {
	h := len(lines)
	if h == 0 {
		return lines
	}
	w := len(lines[0])
	rv := make([][]rune, w, w)

	for i := h - 1; 0 <= i; i-- {
		s := lines[i]
		for j, c := range s {
			if i == h-1 {
				rv[j] = make([]rune, h, h)
			}
			rv[j][h-i-1] = c
		}
	}
	return Map(rv, func(rs []rune, _ int) string {
		return string(rs)
	})
}

// 転置
// transpose
func TransposeLines(lines []string) []string {
	h := len(lines)
	if h == 0 {
		return lines
	}
	w := len(lines[0])
	rv := make([][]rune, w, w)
	for i, row := range lines {
		for j, r := range row {
			if i == 0 {
				rv[j] = make([]rune, h, h)
			}
			rv[j][i] = r
		}
	}
	return Map(rv, func(rs []rune, _ int) string {
		return string(rs)
	})
}
