package piscine

import "fmt"

func Solve(core *Core) {
	fn := firstSize(core)
	if vFlag {
		PrintStr(ScreenClear)
	}
	nMino := len(core.GivenMinos)
	parity := Reduce(core.GivenMinos, 0, func(s, t, _ int) int {
		if MinoTU <= t && t <= MinoTL {
			return s + 1
		}
		return s
	})
	for n := fn; true; n++ {
		nVacant := n*n - nMino*4
		if nVacant == 0 {
			if (parity+n*n)%2 != 0 {
				continue
			}
		}
		fmt.Printf("try for size %d, minos = %d(%c - %c), vacants = %d, parity = %d\n",
			n, nMino,
			Base[0], Base[nMino-1],
			nVacant,
			parity%2,
		)
		if solveForSize(core, n) {
			PrintStr("solved\n")
			return
		}
		PrintStr("failed\n")
	}
}

type SolverState struct {
	// サイズ
	Size int
	// 計算中の盤面(rune)
	Board [][]rune
	// 計算中の盤面(bit)
	BitVBoard, BitHBoard []uint64
	// 位置をkey, ミノ種別をvalueとするマップ
	Placement map[int]int
	// ミノ種別ごとの位置のスタック
	PositionStack map[int]*[]int
	// ミノ種別をkey, 残りの数をvalueとするマップ
	RestMinoCount map[int]int
	FailMap       map[string]bool
}

func solveForSize(core *Core, size int) bool {
	ss := &SolverState{
		Size: size,
		Board: Map(Seq(0, size), func(_, _ int) []rune {
			row := make([]rune, size, size)
			for j := 0; j < size; j++ {
				row[j] = '.'
			}
			return row
		}),
		BitVBoard:     makeBitBoard(size),
		BitHBoard:     makeBitBoard(size),
		Placement:     map[int]int{},
		PositionStack: map[int]*[]int{},
		RestMinoCount: map[int]int{},
		FailMap:       map[string]bool{},
	}
	ForEach(core.Minos, func(k, _ int) {
		s := make([]int, 0, size)
		ss.PositionStack[k] = &s
	})
	ForEach(core.GivenMinos, func(k, _ int) {
		ss.RestMinoCount[k] = ss.RestMinoCount[k] + 1
	})
	return dfs(core, ss, 0)
}

func dfs(core *Core, state *SolverState, k int) bool {
	if k == len(core.GivenMinos) {
		state.bitPrintBoard(core)
		return true
	}

	code := state.encodeBoard()
	if state.FailMap[code] {
		Visualize(core, state, -1, -1, k,
			fmt.Sprintf("HIT!!!!!!!!!!!!!!!!!!!!!(%d,%d)", k, len(state.FailMap)),
		)
		return false
	}

	// if k < len(core.GivenMinos)/2 || k%4 == 3 {
	restVacant := state.getRestVacantPlacableFine(core, k)
	restToPut := 4 * (len(core.GivenMinos) - k)
	// fmt.Println(restVacant, restToPut, residual)
	if restVacant < restToPut {
		// Visualize(core, state, -1, -1, k,
		// 	fmt.Sprintf("no enough vacant: needed %d, actual %d", restToPut, restVacant),
		// )
		// state.FailMap[code] = true
		return false
	}
	// }

	mt := core.GivenMinos[k]
	mino := core.MinoReverseMap[mt]
	// fmt.Println(mino)

	i, j := 0, 0
	// すでに同種のミノを置いていた場合、それよりも後ろに置く必要がある
	if len(*state.PositionStack[mt]) > 0 {
		pos := Top(state.PositionStack[mt]) + 1
		i, j = pos/state.Size, pos%state.Size
	}

	restToPut = state.RestMinoCount[k] * 4
	for ; i < state.Size-mino.Height+1; i++ {
		for ; j < state.Size-mino.Width+1; j++ {
			restVacant := state.getRestVacantOver(i, j)
			if restVacant < restToPut {
				// Visualize(core, state, i, j, k, fmt.Sprintf("no enough vacant for %d", k))
				// state.FailMap[code] = true
				return false
			}

			if !state.isBitPlacableAt(mino, i, j) {
				// Visualize(core, state, i, j, k, fmt.Sprintf("not placable %d at (%d, %d)", k, i, j))
				continue
			}

			state.bitPlaceAt(mino, i, j, k)
			// fmt.Println(state.BitVBoard)
			// Visualize(core, state, i, j, k, fmt.Sprintf("place %d at (%d, %d)", k, i, j))
			if dfs(core, state, k+1) {
				return true
			}
			state.bitRemoveFrom(mino, i, j, k)
		}
		j = 0
	}

	state.FailMap[code] = true
	return false
}
