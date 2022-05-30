package piscine

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type xy struct {
	X int
	Y int
}

func genXY(x int, y int) xy {
	return xy{
		X: x,
		Y: y,
	}
}

func pickXY(arr []xy, index int) ([]xy, xy) {
	return append(arr[:index], arr[index+1:]...), arr[index]
}

func genPosArr(height int, width int) (ret []xy) {
	ret = make([]xy, height*width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ret[y*width+x] = genXY(x, y)
		}
	}
	return
}

func BoardGeneratorStr(heightStr string, widthStr string, blackCount string, numCountStr string) ([][]int, error) {
	height := getRand(heightStr, 2, math.MaxInt)
	width := getRand(widthStr, 2, math.MaxInt)
	// Blackはだいたい1/3くらいが限度。確定値の導出は諦めた。
	black_count := getRand(blackCount, 0, height*width/3)
	num_count := getRand(numCountStr, 0, (height*width)-black_count)

	// ボード生成
	return BoardGenerator(height, width, black_count, num_count)
}

func BoardGenerator(height int, width int, black_count int, num_count int) (result [][]int, err error) {
	rand.Seed(time.Now().UnixNano())

	result = make([][]int, height)
	for y := 0; y < height; y++ {
		result[y] = make([]int, width)
	}

	counter := 0
	blackPos := genPosArr(height, width)
	for i := 0; i < black_count; i++ {
		counter++
		if len(blackPos) <= 0 {
			err = errors.New(fmt.Sprintf("Failed to generate board (Processing: Put Black Piece) (height:%d, width: %d, Black: %d, Num:%d)\n", height, width, black_count, num_count))
			return
		}

		var _xy xy
		blackPos, _xy = pickXY(blackPos, rand.Intn(len(blackPos)))

		x := _xy.X
		y := _xy.Y

		if isThis(result, y, x, GENERATOR_BLACK_NUM) ||
			isThis(result, y, x-1, GENERATOR_BLACK_NUM) ||
			isThis(result, y, x+1, GENERATOR_BLACK_NUM) ||
			isThis(result, y-1, x, GENERATOR_BLACK_NUM) ||
			isThis(result, y+1, x, GENERATOR_BLACK_NUM) {
			i--
			continue
		}

		result[y][x] = GENERATOR_BLACK_NUM

		if containsIsland(result, GENERATOR_BLACK_NUM) {
			i--
			result[y][x] = 0
			continue
		}
	}

	counter = 0
	numPos := genPosArr(height, width)
	for i := 0; i < num_count; i++ {
		if len(numPos) <= 0 {
			err = errors.New(fmt.Sprintf("Failed to generate board (Processing: CalcWhite) (height:%d, width: %d, Black: %d, Num:%d)\n", height, width, black_count, num_count))
			return
		}

		var _xy xy
		numPos, _xy = pickXY(numPos, rand.Intn(len(numPos)))

		x := _xy.X
		y := _xy.Y

		if isThis(result, y, x, GENERATOR_BLACK_NUM) || !isThis(result, y, x, GENERATOR_WHITE_NUM) {
			i--
			continue
		}
		result[y][x] = calcWhite(result, y, x)
	}

	// 導出可能かどうかのチェックは後で実装する
	return
}

func isThis(target [][]int, y int, x int, val int) bool {
	if y < 0 || len(target) <= y || x < 0 || len(target[0]) <= x {
		return false
	}
	return target[y][x] == val
}

func calcWhite(board [][]int, y int, x int) int {
	retnum := 0

	for _x := x - 1; _x >= -1; _x-- {
		if _x < 0 || board[y][_x] == GENERATOR_BLACK_NUM {
			break
		}
		retnum++
	}

	for _y := y - 1; _y >= -1; _y-- {
		if _y < 0 || board[_y][x] == GENERATOR_BLACK_NUM {
			break
		}
		retnum++
	}

	for _x := x + 1; _x <= len(board[0]); _x++ {
		if _x == len(board[0]) || board[y][_x] == GENERATOR_BLACK_NUM {
			break
		}
		retnum++
	}

	for _y := y + 1; _y <= len(board); _y++ {
		if _y == len(board) || board[_y][x] == GENERATOR_BLACK_NUM {
			break
		}
		retnum++
	}

	return retnum + 1
}
