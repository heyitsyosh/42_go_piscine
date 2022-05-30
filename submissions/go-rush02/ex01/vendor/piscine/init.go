package piscine

import "flag"

var (
	vFlag = false
	GFlag = 0
	speedFlag = 300
)

func init() {	
	flag.BoolVar(&vFlag, "v", false, "enable visualizer.")
	flag.IntVar(&GFlag, "g", 0, "generate random fillit file. -g=MINO")
	flag.IntVar(&speedFlag, "s", 300, "specify the visualize speed. -s=SPEED")
	flag.Parse()
}