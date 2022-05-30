package piscine

type bsq struct {
	empty, obstacle, full rune
	max_x, max_y          int
	x, y, area            int
}

func First_line(input string) bool {
	b := bsq{}
	//take \n index
	r_input := []rune(input)
	var i int
	for j, r := range r_input {
		if r == '\n' {
			i = j
			break
		}
	}
	if len(r_input[:i]) < 4 {
		PrintError("Invalid first line")
		return false
	}
	//check validity of specified amount of lines(rows), atoi
	rowCount, isValidSide := Atoi(string(r_input[:i-3]))
	if !isValidSide {
		PrintError("Invalid line specification")
		return false
	}
	b.max_x = rowCount
	//check validity of specified empty, obstacle, fill runes
	runes := r_input[i-3 : i]
	for j := 0; j < 2; j++ {
		if runes[j] == runes[j+1] {
			PrintError("Invalid fill, empty, obstacle specification")
			return false
		}
	}
	if runes[0] == runes[2] {
		PrintError("Invalid fill, empty, obstacle specification")
		return false
	}

	b.empty, b.obstacle, b.full = runes[0], runes[1], runes[2]
	setUpMap, status := map_info(r_input[i+1:], &b)
	if !status {
		return false
	}
	solved := Solve(setUpMap, &b)
	Print_output(solved, &b)
	return true
}

func map_info(r_input []rune, b *bsq) ([][]int, bool) {
	//find width of first line
	var i int
	var index int
	for j, r := range r_input {
		if r == '\n' {
			i = j
			break
		}
	}

	b.max_y = len(r_input[:i])
	setUpMap := make([][]int, b.max_x)
	for k := range setUpMap {
		setUpMap[k] = make([]int, b.max_y)
	}
	if b.max_y <= 0 {
		PrintError("Invalid map")
		return setUpMap, false
	}
	if (b.max_x)*(b.max_y+1) != len(r_input) {
		PrintError("Expected map length and actual map length mismatch")
		return setUpMap, false
	}
	//initialize int array; empty = 1, obstacle = 0;
	//make sure all lengths are equal, that only empty and obstacle runes are in string
	for p := 0; p < b.max_x && index < len(r_input); p++ {
		for q := 0; q < b.max_y + 1 && index < len(r_input); q++ {
			if r_input[index] == b.empty {
				setUpMap[p][q] = 1
			} else if r_input[index] == b.obstacle {
				setUpMap[p][q] = 0
			} else if r_input[index] == '\n' && q == b.max_y {
				index++
				continue
			} else {
				PrintError("Invalid map") //doesn't have \n at end or contains unspecified characters
				return setUpMap, false
			}
			index++
		}
	}
	return setUpMap, true
}
