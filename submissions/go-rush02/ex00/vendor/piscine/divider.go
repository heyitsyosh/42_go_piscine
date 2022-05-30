package piscine

// ファイルの中身をテトリミノごとのブロック([]string)に分割する。
func Divide(content string) ([][]string, bool) {
	blocks := Map(Split(content, "\n\n"), func(s string, _ int) []string {
		return Split(StringTrim(s, "\n"), "\n")
	})
	// fmt.Println(blocks)
	return blocks, true
}
