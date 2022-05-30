package piscine

import "fmt"

// 1 <= size <= 64
func makeBitBoard(size int) []uint64 {
	return Map(Seq(0, size), func(_, _ int) uint64 {
		return 0
	})
}

func (ss *SolverState) bitPrintBoard(core *Core) {
	board := ss.makeBitBoardString(core)
	ForEach(board, func(rs []rune, _ int) {
		fmt.Println(string(rs))
	})
}

func (ss *SolverState) makeBitBoardString(core *Core) [][]rune {
	board := Map(Seq(0, ss.Size), func(_, _ int) []rune {
		row := make([]rune, ss.Size)
		for i, _ := range Seq(0, ss.Size) {
			row[i] = '.'
		}
		return row
	})
	for pos, k := range ss.Placement {
		i, j := pos/ss.Size, pos%ss.Size
		t := core.GivenMinos[k]
		mino := core.MinoReverseMap[t]
		r := rune(Base[k])
		for ii, row := range mino.shape {
			for jj, c := range row {
				if c != '#' {
					continue
				}
				board[i+ii][j+jj] = r
			}
		}
	}
	return board
}

func (ss *SolverState) isBitPlacableAt(mino *MinoMaster, i, j int) bool {
	if len(mino.bitVShape) <= len(mino.bitHShape) {
		for ii, bitRow := range mino.bitVShape {
			mask := (bitRow << j)
			if (ss.BitVBoard[i+ii] & mask) != 0 {
				return false
			}
		}
	} else {
		for jj, bitCol := range mino.bitHShape {
			mask := (bitCol << i)
			if (ss.BitHBoard[j+jj] & mask) != 0 {
				return false
			}
		}
	}
	return true
}

func (ss *SolverState) bitPlaceAt(mino *MinoMaster, i, j, k int) {
	ss.Placement[i*ss.Size+j] = k
	Push(ss.PositionStack[mino.MinoType], i*ss.Size+j)
	ss.RestMinoCount[mino.MinoType] = ss.RestMinoCount[mino.MinoType] - 1
	for ii, bitRow := range mino.bitVShape {
		ss.BitVBoard[i+ii] |= (bitRow << j)
	}
	for jj, bitCol := range mino.bitHShape {
		ss.BitHBoard[j+jj] |= (bitCol << i)
	}
}

func (ss *SolverState) bitRemoveFrom(mino *MinoMaster, i, j, k int) {
	delete(ss.Placement, i*ss.Size+j)
	Pop(ss.PositionStack[mino.MinoType])
	ss.RestMinoCount[mino.MinoType] = ss.RestMinoCount[mino.MinoType] + 1
	for ii, bitRow := range mino.bitVShape {
		ss.BitVBoard[i+ii] ^= (bitRow << j)
	}
	for jj, bitCol := range mino.bitHShape {
		ss.BitHBoard[j+jj] ^= (bitCol << i)
	}
}

func (ss *SolverState) getRestVacantOver(i, j int) int {
	n := 0
	for ii := i; ii < ss.Size; ii++ {
		var mask uint64 = (1 << ss.Size) - 1
		if ii == i {
			mask = (1 << ss.Size) - (1 << j)
		}
		n += PopCount((^ss.BitVBoard[ii]) & mask)
		// fmt.Println(ii, mask, n)
	}
	return n
}

func (ss *SolverState) fill(bb []uint64, i, j int, f func(i, j int)) int {
	bb[i] |= (1 << j)
	if ((ss.BitVBoard[i] >> j) & 1) == 1 {
		return 0
	}
	if f != nil {
		f(i, j)
	}
	var n int = 1
	if 0 < i && ((bb[i-1]>>j)&1) == 0 {
		n += ss.fill(bb, i-1, j, f)
	}
	if i+1 < ss.Size && ((bb[i+1]>>j)&1) == 0 {
		n += ss.fill(bb, i+1, j, f)
	}
	if 0 < j && ((bb[i]>>(j-1))&1) == 0 {
		n += ss.fill(bb, i, j-1, f)
	}
	if j+1 < ss.Size && ((bb[i]>>(j+1))&1) == 0 {
		n += ss.fill(bb, i, j+1, f)
	}
	return n
}

// 白マスによる「島」について「サイズを4で割ったあまりを除外したもの」の総和を返す
// -> 「残りの白マスのうちテトラミノを使って埋められる可能性があるマスの数」の上からの評価になる
func (ss *SolverState) getRestVacantPlacable() (int, []int, map[int]int) {
	bb := Map(ss.BitVBoard, func(s uint64, _ int) uint64 {
		return s
	})
	var n int = 0
	isle4 := make([]int, 0, ss.Size*ss.Size)
	over4 := map[int]int{}
	for i := 0; i < ss.Size; i++ {
		for j := 0; j < ss.Size; j++ {
			if ((bb[i] >> j) & 1) == 1 {
				continue
			}
			nn := ss.fill(bb, i, j, nil)
			pos := i*ss.Size + j
			if nn == 4 {
				isle4 = append(isle4, pos)
			} else if 4 < nn && nn < 8 {
				over4[pos] = nn
			} else if nn == 8 {
				over4[pos] = nn
			}
			n += nn - nn%4
		}
	}
	return n, isle4, over4
}

func (ss *SolverState) getRestVacantPlacableFine(core *Core, k int) int {
	n, isle4, _ := ss.getRestVacantPlacable()

	// サイズ4の島について、残りの手駒で埋められるかどうかをチェックする。
	restMino := ss.RestMinoCount
	for _, pos := range isle4 {
		i, j := pos/ss.Size, pos%ss.Size
		ok := false
		for kk, count := range restMino {
			if count <= 0 {
				continue
			}
			mino := core.MinoReverseMap[kk]
			j -= mino.firstBlack
			if j < 0 {
				continue
			}
			if ss.Size < i+mino.Height || ss.Size < j+mino.Width {
				continue
			}
			if ss.isBitPlacableAt(mino, i, j) {
				restMino[kk] -= 1
				ok = true
				break
			}
		}
		if !ok {
			// 埋められない島があるなら、その分4マスをnから引く
			n -= 4
		}
	}
	return n
}

const Base = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_@"

func (ss *SolverState) encodeBoard() string {
	var s = ""
	for i := 0; i < ss.Size; i++ {
		row := ss.BitVBoard[i]
		for j := 0; j < ss.Size; j += 6 {
			x := Base[(row>>j)&0x3f]
			s += string(x)
		}
	}
	return s
}
