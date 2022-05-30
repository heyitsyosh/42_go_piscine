package piscine

func Symbolize(core *Core, blocks [][]string) []int {
	return Map(blocks, func(lines []string, _ int) int {
		trimmed := TrimBlankRowCol(lines)
		joined := Join(trimmed, "\n")
		mino, _ := core.MinoMap[joined]
		return mino.MinoType
	})
}
