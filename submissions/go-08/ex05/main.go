package main

import "ft"

const (
	OPEN = 1
	CLOSE = 2
)

type Door struct {
	state int
}

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) {
	if ptrDoor == nil {
		return
	}
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	if ptrDoor == nil {
		return
	}
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(door Door) bool {
	PrintStr("is the Door opened ?")
	return door.state == OPEN
}

func IsDoorClose(door Door) bool {
	PrintStr("is the Door closed ?")
	return door.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(*door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}