package robot_room_cleaner

import "testing"

type point struct {
	r, c int
}

type mockRobot struct {
	room    [][]int
	row     int
	col     int
	dir     int
	cleaned map[point]bool
	dr      []int
	dc      []int
}

func newMockRobot(room [][]int, row, col int) *mockRobot {
	return &mockRobot{
		room:    room,
		row:     row,
		col:     col,
		dir:     0,
		cleaned: make(map[point]bool),
		dr:      []int{-1, 0, 1, 0},
		dc:      []int{0, 1, 0, -1},
	}
}

func (r *mockRobot) Move() bool {
	nr := r.row + r.dr[r.dir]
	nc := r.col + r.dc[r.dir]
	if nr < 0 || nr >= len(r.room) || nc < 0 || nc >= len(r.room[0]) || r.room[nr][nc] == 0 {
		return false
	}
	r.row = nr
	r.col = nc
	return true
}

func (r *mockRobot) TurnLeft() {
	r.dir = (r.dir + 3) % 4
}

func (r *mockRobot) TurnRight() {
	r.dir = (r.dir + 1) % 4
}

func (r *mockRobot) Clean() {
	r.cleaned[point{r.row, r.col}] = true
}

func TestCleanRoom(t *testing.T) {
	tests := []struct {
		name     string
		room     [][]int
		startRow int
		startCol int
	}{
		{
			name: "simple grid",
			room: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			startRow: 0,
			startCol: 0,
		},
		{
			name: "grid with walls",
			room: [][]int{
				{1, 1, 0},
				{1, 1, 1},
				{0, 1, 1},
			},
			startRow: 1,
			startCol: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			robot := newMockRobot(tt.room, tt.startRow, tt.startCol)
			cleanRoom(robot)

			expected := 0
			for r := 0; r < len(tt.room); r++ {
				for c := 0; c < len(tt.room[0]); c++ {
					if tt.room[r][c] == 1 {
						expected++
						if !robot.cleaned[point{r, c}] {
							t.Errorf("cell (%d,%d) was not cleaned", r, c)
						}
					}
				}
			}

			if len(robot.cleaned) != expected {
				t.Errorf("cleaned %d cells, want %d", len(robot.cleaned), expected)
			}
		})
	}
}
