package piscine

func isValidInput(input []string) bool {
	// 空入力チェック
	height := len(input)
	if height <= 0 {
		return false
	}
	width := len(input[0])
	if width <= 0 {
		return false
	}
	if height != width {
		return false
	}
	for _, v := range input {
		if width != len(v) {
			// 幅が揃っているかチェック
			return false
		}

		// 変な文字が混入していないかチェック
		// 幅とかが正しいかチェックはParse部分で行う
		for _, c := range v {
			// 0と1は無効な入力 (0はBlack相当, 1は島ができるためNG)
			if !(('2' <= c && c <= '9') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '.') {
				return false
			}
		}
	}
	return true
}
