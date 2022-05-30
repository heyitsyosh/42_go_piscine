package piscine

type Core struct {
	// ミノ種別のマスターデータ
	Minos []int
	// ミノ形状のマスターデータ
	MinoMap        map[string]*MinoMaster
	MinoReverseMap map[int]*MinoMaster
	// 入力をミノ種別に変換したもの
	GivenMinos []int
}

func NewCore() *Core {
	MinoMap, MinoReverseMap := MakeMinoShapes()
	return &Core{
		Minos: []int{
			MinoO,
			MinoLU,
			MinoLR,
			MinoLD,
			MinoLL,
			MinoJU,
			MinoJR,
			MinoJD,
			MinoJL,
			MinoTU,
			MinoTR,
			MinoTD,
			MinoTL,
			MinoIU,
			MinoIR,
			MinoSU,
			MinoSR,
			MinoZU,
			MinoZR,
		},
		MinoMap:        MinoMap,
		MinoReverseMap: MinoReverseMap,
	}
}
