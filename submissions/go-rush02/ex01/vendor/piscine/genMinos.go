package piscine

// width = len(str[0]])
// hight = len(str)
// 0 - (4 - width)
// 0 - (4 - height)

func Minos(minoType int) []string {
	var str []string
	switch minoType {
	case MinoO:
		str = append(str,
			"##",
			"##")
	case MinoLU:
		str = append(str,
			"#.",
			"#.",
			"##")
	case MinoLR:
		str = append(str,
			"###",
			"#..")
	case MinoLD:
		str = append(str,
			"##",
			".#",
			".#")
	case MinoLL:
		str = append(str,
			"..#",
			"###")
	case MinoJU:
		str = append(str,
			".#",
			".#",
			"##")
	case MinoJR:
		str = append(str,
			"#..",
			"###")
	case MinoJD:
		str = append(str,
			"##",
			"#.",
			"#.")
	case MinoJL:
		str = append(str,
			"###",
			"#..")
	case MinoTU:
		str = append(str,
			"###",
			".#.")
	case MinoTR:
		str = append(str,
			".#",
			"##",
			".#")
	case MinoTD:
		str = append(str,
			".#.",
			"###")
	case MinoTL:
		str = append(str,
			"#.",
			"##",
			"#.")
	case MinoIU:
		str = append(str,
			"#",
			"#",
			"#",
			"#")
	case MinoIR:
		str = append(str,
			"####")
	case MinoSU:
		str = append(str,
			".##",
			"##.")
	case MinoSR:
		str = append(str,
			"#.",
			"##",
			".#")
	case MinoZU:
		str = append(str,
			"##.",
			".##")
	case MinoZR:
		str = append(str,
			".#",
			"##",
			"#.")
	}
	return str
}

func TransMinos(w,h int, mino []string) []string {
	r := 4 - w - len(mino[0])
	for i := range mino {
		for j := 0; j < w; j++ {
			mino[i] = "." + mino[i]
		}
		for j := 0; j < r; j++ {
			mino[i] += "."
		}
		mino[i] += "\n"
	}
	var tmp []string
	r = 4 - h - len(mino)
	for i := 0; i < h; i++ {
		tmp = append(tmp, "....\n")
	}
	for i := range mino {
		tmp = append(tmp, mino[i])
	}
	for i := 0; i < r; i++ {
		tmp = append(tmp, "....\n")
	}
	tmp[3] = tmp[3][0:4]
	return tmp
}