package robot_room_cleaner

type Robot interface {
	Move() bool
	TurnLeft()
	TurnRight()
	Clean()
}

func cleanRoom(robot Robot) {
}
