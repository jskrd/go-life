package models_test

import (
	"testing"

	"github.com/jskrd/go-life/internal/models"
)

func TestIsAlive(t *testing.T) {
	t.Run("empty generation returns false", func(t *testing.T) {
		generation := models.Generation{Cells: make(map[[2]int]struct{})}

		expected := false
		actual := generation.IsAlive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns true for live cell", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{{0, 0}: {}}}

		expected := true
		actual := generation.IsAlive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns false for dead cell", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{{0, 0}: {}}}

		expected := false
		actual := generation.IsAlive(1, 1)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})
}

func TestSetAlive(t *testing.T) {
	t.Run("sets a cell alive", func(t *testing.T) {
		generation := models.Generation{Cells: make(map[[2]int]struct{})}
		generation.SetAlive(0, 0)

		_, exists := generation.Cells[[2]int{0, 0}]
		if !exists {
			t.Error("expected cell to be alive but it was dead")
		}
	})

	t.Run("sets multiple cells alive", func(t *testing.T) {
		generation := models.Generation{Cells: make(map[[2]int]struct{})}
		generation.SetAlive(0, 0)
		generation.SetAlive(1, 1)

		expected := true
		actual1 := generation.IsAlive(0, 0)
		actual2 := generation.IsAlive(1, 1)
		if actual1 != expected || actual2 != expected {
			t.Errorf("expected both cells to be %v but got %v and %v", expected, actual1, actual2)
		}
	})

	t.Run("sets an already live cell again", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{{0, 0}: {}}}
		generation.SetAlive(0, 0)

		_, exists := generation.Cells[[2]int{0, 0}]
		if !exists {
			t.Error("expected cell to be alive but it was dead")
		}
	})
}

func TestCountNeighbors(t *testing.T) {
	t.Run("returns zero for an empty generation", func(t *testing.T) {
		generation := models.Generation{Cells: make(map[[2]int]struct{})}

		expected := uint(0)
		actual := generation.CountNeighbors(0, 0)
		if actual != expected {
			t.Errorf("expected %d neighbors but got %d", expected, actual)
		}
	})

	t.Run("returns zero for a generation with one live cell", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{{0, 0}: {}}}

		expected := uint(0)
		actual := generation.CountNeighbors(0, 0)
		if actual != expected {
			t.Errorf("expected %d neighbors but got %d", expected, actual)
		}
	})

	t.Run("returns one for a generation with one live neighbor", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{0, 0}: {},
			{1, 1}: {},
		}}

		expected := uint(1)
		actual := generation.CountNeighbors(0, 0)
		if actual != expected {
			t.Errorf("expected %d neighbors but got %d", expected, actual)
		}
	})

	t.Run("returns two for a generation with two live neighbors", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{0, 0}: {},
			{1, 1}: {},
			{1, 0}: {},
		}}

		expected := uint(2)
		actual := generation.CountNeighbors(0, 0)
		if actual != expected {
			t.Errorf("expected %d neighbors but got %d", expected, actual)
		}
	})

	t.Run("returns eight for a generation with eight live neighbors", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{-1, -1}: {}, {0, -1}: {}, {1, -1}: {},
			{-1, 0}: {}, {0, 0}: {}, {1, 0}: {},
			{-1, 1}: {}, {0, 1}: {}, {1, 1}: {},
		}}

		expected := uint(8)
		actual := generation.CountNeighbors(0, 0)
		if actual != expected {
			t.Errorf("expected %d neighbors but got %d", expected, actual)
		}
	})

	t.Run("returns one for each possible neighbor position", func(t *testing.T) {
		neighborPositions := [][2]int{
			{-1, -1},
			{0, -1},
			{1, -1},
			{1, 0},
			{1, 1},
			{0, 1},
			{-1, 1},
			{-1, 0},
		}

		for _, pos := range neighborPositions {
			generation := models.Generation{Cells: make(map[[2]int]struct{})}
			generation.SetAlive(0, 0)
			generation.SetAlive(pos[0], pos[1])

			expected := uint(1)
			actual := generation.CountNeighbors(0, 0)
			if actual != expected {
				t.Errorf("for neighbor at (%d, %d): expected %d neighbors but got %d",
					pos[0], pos[1], expected, actual)
			}
		}
	})

	t.Run("returns zero for a generation with one live cell and one live non-neighbor", func(t *testing.T) {
		nonNeighborPositions := [][2]int{
			{-2, -2},
			{-1, -2},
			{0, -2},
			{1, -2},
			{2, -2},
			{2, -1},
			{2, 0},
			{2, 1},
			{2, 2},
			{1, 2},
			{0, 2},
			{-1, 2},
			{-2, 2},
			{-2, 1},
			{-2, 0},
			{-2, -1},
		}

		for _, pos := range nonNeighborPositions {
			generation := models.Generation{Cells: make(map[[2]int]struct{})}
			generation.SetAlive(0, 0)
			generation.SetAlive(pos[0], pos[1])

			expected := uint(0)
			actual := generation.CountNeighbors(0, 0)
			if actual != expected {
				t.Errorf("for non-neighbor at (%d, %d): expected %d neighbors but got %d",
					pos[0], pos[1], expected, actual)
			}
		}
	})
}

func TestShouldLive(t *testing.T) {
	t.Run("returns false for a dead cell with zero live neighbors", func(t *testing.T) {
		generation := models.Generation{Cells: make(map[[2]int]struct{})}

		expected := false
		actual := generation.ShouldLive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns false for a dead cell with one live neighbor", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{{1, 1}: {}}}

		expected := false
		actual := generation.ShouldLive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns false for a dead cell with two live neighbors", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{0, -1}: {},
			{1, -1}: {},
		}}

		expected := false
		actual := generation.ShouldLive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns true for a dead cell with three live neighbors", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{-1, -1}: {},
			{0, -1}:  {},
			{1, -1}:  {},
		}}

		expected := true
		actual := generation.ShouldLive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns false for a dead cell with four live neighbors", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{-1, -1}: {},
			{0, -1}:  {},
			{1, -1}:  {},
			{-1, 0}:  {},
		}}

		expected := false
		actual := generation.ShouldLive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns false for a live cell with one live neighbor", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{0, 0}: {},
			{1, 1}: {},
		}}

		expected := false
		actual := generation.ShouldLive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns true for a live cell with two live neighbors", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{0, 0}: {},
			{1, 1}: {},
			{1, 0}: {},
		}}

		expected := true
		actual := generation.ShouldLive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns true for a live cell with three live neighbors", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{-1, -1}: {},
			{0, -1}:  {},
			{1, -1}:  {},
			{0, 0}:   {},
		}}

		expected := true
		actual := generation.ShouldLive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})

	t.Run("returns false for a live cell with four live neighbors", func(t *testing.T) {
		generation := models.Generation{Cells: map[[2]int]struct{}{
			{-1, -1}: {},
			{0, -1}:  {},
			{1, -1}:  {},
			{-1, 0}:  {},
			{0, 0}:   {},
		}}

		expected := false
		actual := generation.ShouldLive(0, 0)
		if actual != expected {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})
}
