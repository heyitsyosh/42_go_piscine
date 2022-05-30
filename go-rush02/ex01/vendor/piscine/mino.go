package piscine

// ミノタイプ
const MinoTypes = 19
const (
	MinoO = iota
	MinoLU
	MinoLR
	MinoLD
	MinoLL
	MinoJU
	MinoJR
	MinoJD
	MinoJL
	MinoTU
	MinoTR
	MinoTD
	MinoTL
	MinoIU
	MinoIR
	MinoSU
	MinoSR
	MinoZU
	MinoZR
)

var MinoNames []string = []string{
	"O",
	"LU",
	"LR",
	"LD",
	"LL",
	"JU",
	"JR",
	"JD",
	"JL",
	"TU",
	"TR",
	"TD",
	"TL",
	"IU",
	"IR",
	"SU",
	"SR",
	"ZU",
	"ZR",
}

const MinoSize = 4

type MinoMaster struct {
	MinoType             int
	Height, Width        int
	firstBlack           int // 一番上の行の一番左の'#'の横位置
	shape                [][]rune
	bitHShape, bitVShape []uint64
}

func MakeMinoShapes() (map[string]*MinoMaster, map[int]*MinoMaster) {
	shapes := make([][]string, 0)
	rotator := func(shape []string, n int) {
		s := shape
		for k := 0; k < n; k++ {
			shapes = append(shapes, s)
			s = RotateLines(s)
		}
	}

	shapeO := []string{
		"##",
		"##",
	}
	shapeL := []string{
		"#.",
		"#.",
		"##",
	}
	// J = mirror of L
	shapeJ := ReverseLines(shapeL)
	shapeT := []string{
		"###",
		".#.",
	}
	shapeS := []string{
		".##",
		"##.",
	}
	// Z = mirror of S
	shapeZ := ReverseLines(shapeS)
	shapeI := []string{
		"#",
		"#",
		"#",
		"#",
	}

	rotator(shapeO, 1)
	rotator(shapeL, 4)
	rotator(shapeJ, 4)
	rotator(shapeT, 4)
	rotator(shapeI, 2)
	rotator(shapeS, 2)
	rotator(shapeZ, 2)

	shapeMap := map[string]*MinoMaster{}
	shapeReverseMap := map[int]*MinoMaster{}
	for i, shape := range shapes {
		h := len(shape)
		w := len(shape[0])
		joined := Join(shape, "\n")
		runeShape := Map(shape, func(s string, _ int) []rune {
			return []rune(s)
		})
		firstBlack := Index([]rune(shape[0]), func(r rune, _ int) bool {
			return r == '#'
		})
		mm := &MinoMaster{
			MinoType:   i,
			Height:     h,
			Width:      w,
			shape:      runeShape,
			firstBlack: firstBlack,
			bitVShape:  ShapeToBitmask(shape),
			bitHShape:  ShapeToBitmask(TransposeLines(shape)),
		}
		shapeMap[joined] = mm
		shapeReverseMap[i] = mm
		// fmt.Println(mm)
	}
	return shapeMap, shapeReverseMap
}

func ShapeToBitmask(shape []string) []uint64 {
	return Map(shape, func(s string, _ int) uint64 {
		var v uint64
		for i, c := range s {
			if c == '#' {
				v |= (1 << i)
			}
		}
		return v
	})
}
