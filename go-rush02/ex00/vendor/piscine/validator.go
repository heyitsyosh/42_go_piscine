package piscine

// import "fmt"

// blocksの各要素([]string)が正しいテトリミノ定義であることを確認する。
// check that every elements in block is valid as tetrimino definition.
func Validate(core *Core, blocks [][]string) bool {
	n := len(blocks)
	if n < 1 || len(Base) < n {
		PrintError("number of block is invalid")
		// fmt.Println("n", n)
		return false
	}

	return Every(blocks, func(lines []string, i int) bool {
		// []stringの長さが4であるか？
		// size of lines is 4?
		h := SliceLen(lines)
		if h != MinoSize {
			PrintError("mino height is invalid")
			// fmt.Println("-------")
			// fmt.Println(Join(lines, "\n"), i)
			// fmt.Println("-------")
			return false
		}
		// 各要素の長さが4であるか？
		// size of every rows are 4?
		ok := Every(lines, func(line string, _ int) bool {
			w := SliceLen([]rune(line))
			return w == MinoSize
		})
		if !ok {
			PrintError("mino width is invalid")
			return false
		}

		// 3. すべての文字が`#`か`.`であるか
		// every characters are '#' or '.'?
		runeCount := map[rune]int{}
		for _, r := range Join(lines, "") {
			runeCount[r] = runeCount[r] + 1
		}
		hasInvalidChar := runeCount['#']+runeCount['.'] != MinoSize*MinoSize
		if hasInvalidChar {
			PrintError("mino contains unexpected character")
			return false
		}
		// 3. `#`がちょうど4つあるか？
		// `lines`` contains just 4 '#'?
		sharps := runeCount['#']
		if sharps != 4 {
			PrintError("mino contains more or less numbers '#'")
			return false
		}

		// 4. `#`の配置がテトリミノとして正しいこと
		// '#' arrangement is valid as a tetrimino?
		trimmed := TrimBlankRowCol(lines)
		joined := Join(trimmed, "\n")
		// joinedは標準形になっている
		_, exists := core.MinoMap[joined]
		// fmt.Println(minoType, exists)
		// fmt.Println(joined)
		if !exists {
			PrintError("mino has unexpected shape")
			return false
		}
		return true
	})
}
