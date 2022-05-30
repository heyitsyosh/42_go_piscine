package piscine

func (ss *SolverState) PrintBoard() {
	ForEach(ss.Board, func(rs []rune, _ int) {
		PrintStr(string(rs) + "\n")
	})
}

func (ss *SolverState) isPlacableAt(mino *MinoMaster, i, j int) bool {
	for ii, row := range mino.shape {
		for jj, c := range row {
			if c != '#' {
				continue
			}
			if ss.Board[i+ii][j+jj] != '.' {
				return false
			}
		}
	}
	return true
}

func (ss *SolverState) placeAt(mino *MinoMaster, i, j, k int) {
	r := 'A' + rune(k)
	for ii, row := range mino.shape {
		for jj, c := range row {
			if c != '#' {
				continue
			}
			ss.Board[i+ii][j+jj] = r
		}
	}
}

func (ss *SolverState) removeFrom(mino *MinoMaster, i, j int) {
	for ii, row := range mino.shape {
		for jj, c := range row {
			if c != '#' {
				continue
			}
			ss.Board[i+ii][j+jj] = '.'
		}
	}
}

func firstSize(core *Core) int {
	n := len(core.GivenMinos)
	minos := n * 4
	s := 4
	for ; true; s++ {
		if minos <= s*s {
			break
		}
	}
	return s
}
