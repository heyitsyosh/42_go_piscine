package piscine

import (
	"fmt"
	"math/rand"
	s "strings"
	"time"
)

// n個のテトリミノをランダムに選択し、それを1つの文字列に連結して返す
func Generate(core *Core, n int) (input string) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		randMino := 0 + rand.Intn(19)
		mino := Minos(randMino)
		width, height := len(mino[0]), len(mino)
		randWidth := rand.Intn(5 - width)
		randHeight := rand.Intn(5 - height)
		if i != 0 && i != n {
			input += "\n\n"
		}
		input += s.Join(TransMinos(randWidth, randHeight, mino), "")
	}
	fmt.Println(input)
	return
}
