package universe

import (
	"testing"

	"github.com/jskrd/go-life/internal/models/generation"
)

func TestTick(t *testing.T) {
	tests := []struct {
		name        string
		generations [][][2]int
	}{
		{
			name: "block pattern remains static",
			generations: [][][2]int{
				{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
				{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			},
		},
		{
			name: "blinker pattern oscillates",
			generations: [][][2]int{
				{{1, 2}, {2, 2}, {3, 2}},
				{{2, 1}, {2, 2}, {2, 3}},
				{{1, 2}, {2, 2}, {3, 2}},
			},
		},
		{
			name: "glider pattern moves diagonally",
			generations: [][][2]int{
				{{2, 1}, {3, 2}, {1, 3}, {2, 3}, {3, 3}},
				{{1, 2}, {3, 2}, {2, 3}, {3, 3}, {2, 4}},
				{{3, 2}, {1, 3}, {3, 3}, {2, 4}, {3, 4}},
				{{2, 2}, {3, 3}, {4, 3}, {2, 4}, {3, 4}},
				{{3, 2}, {4, 3}, {2, 4}, {3, 4}, {4, 4}},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			generation := generation.Generation{Cells: make(map[[2]int]struct{})}
			for _, coord := range test.generations[0] {
				generation.SetAlive(coord[0], coord[1])
			}

			universe := Universe{Generation: generation}

			for i := 1; i < len(test.generations); i++ {
				universe.Tick()

				expected := len(test.generations[i])
				actual := len(universe.Generation.Cells)
				if actual != expected {
					t.Errorf("generation %d: expected %d live cells but got %d", i, expected, actual)
				}

				for _, coord := range test.generations[i] {
					if !universe.Generation.IsAlive(coord[0], coord[1]) {
						t.Errorf("generation %d: expected cell at (%d, %d) to be alive but it was dead",
							i, coord[0], coord[1])
					}
				}
			}
		})
	}
}
